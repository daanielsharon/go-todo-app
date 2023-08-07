package setup

import (
	"database/sql"
	"fmt"
	"server/app"
	"server/helper"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
}

func NewTestSetup() TestSetup {
	now := time.Now()
	roleName := fmt.Sprintf("test%v%v", now.Minute(), now.Second())
	password := fmt.Sprintf("Pa$w0rd%v%v", now.Minute(), now.Second())
	return &Setup{
		db:       app.NewDatabase("todoapp_test"),
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
}

func (s *Setup) Close() {
	defer s.db.Close()
	_, err := s.db.Exec(fmt.Sprintf("DROP SCHEMA %v CASCADE;", s.roleName))
	_, err = s.db.Exec(fmt.Sprintf("DROP ROLE %v;", s.roleName))
	helper.PanicIfError(err)
	TruncateAll(s.db)
}
