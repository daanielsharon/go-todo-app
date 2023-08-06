package setup

import (
	"database/sql"
	"server/app"
	"server/helper"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DB() *sql.DB {
	dataSourceName := "postgresql://root:root@localhost:1234/todoapp_test?sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

func Router(db *sql.DB) *gin.Engine {
	validator := validator.New()
	timeout := time.Duration(1) * time.Second

	userController, containerController, itemController := app.NewAppSetup(db, validator, timeout)

	router := app.NewRouter(containerController, itemController, userController)
	return router
}

func All() (*gin.Engine, *sql.DB) {
	db := DB()
	TruncateAll(db)
	router := Router(db)

	return router, db
}
