package setup

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"path/filepath"
	"runtime"
	"server/app"
	"server/helper"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Setup struct {
	db       *sql.DB
	roleName string
	password string
}

type TestSetup interface {
	Router() *gin.Engine
	Open()
	Close()
	Wait() *sync.WaitGroup
	TruncateAll()
}

func NewTestSetup() TestSetup {
	now := time.Now()
	num, err := rand.Int(rand.Reader, big.NewInt(10000000))
	helper.PanicIfError(err)
	roleName := fmt.Sprintf("test%v%v%v", now.Minute(), now.Second(), num)
	password := fmt.Sprintf("Pa$w0rd%v%v%v", now.Minute(), now.Second(), num)
	return &Setup{
		db:       NewTestDatabase("root", "root"),
		roleName: roleName,
		password: password,
	}
}

func (s *Setup) Wait() *sync.WaitGroup {
	return &sync.WaitGroup{}
}

func (s *Setup) Router() *gin.Engine {
	validator := validator.New()
	timeout := time.Duration(1) * time.Second

	userController, containerController, itemController := app.NewAppSetup(s.db, validator, timeout)

	router := app.NewRouter(containerController, itemController, userController)
	return router
}

func (s *Setup) Open() {
	_, err := s.db.Exec(fmt.Sprintf("CREATE ROLE %v WITH LOGIN PASSWORD '%v'", s.roleName, s.password))
	helper.PanicIfError(err)
	_, err = s.db.Exec(fmt.Sprintf("CREATE SCHEMA %v AUTHORIZATION %v;", s.roleName, s.roleName))
	helper.PanicIfError(err)

	// disconnect from db
	s.db.Close()

	// connect to the new schema
	s.db = NewTestDatabase(s.roleName, s.password)

	// migrate to the schema
	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	helper.PanicIfError(err)

	_, basePath, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(basePath), "../../")

	migrate, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%v/migrations", rootDir), "postgres", driver)
	helper.PanicIfError(err)
	migrate.Up()
}

func (s *Setup) Close() {
	// close schema connection
	s.db.Close()

	// back to original connection to drop schema and role
	s.db = NewTestDatabase("root", "root")
	_, err := s.db.Exec(fmt.Sprintf("DROP SCHEMA %v CASCADE;", s.roleName))
	helper.PanicIfError(err)
	_, err = s.db.Exec(fmt.Sprintf("DROP ROLE %v;", s.roleName))
	helper.PanicIfError(err)

	// close connection when everything's done
	defer s.db.Close()
	s.TruncateAll()
}
