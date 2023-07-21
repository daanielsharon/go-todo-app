package setup

import (
	"database/sql"
	"server/app"
	"server/controller"
	"server/helper"
	"server/repository"
	"server/service"
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

	// jwt
	jwtService := service.NewJWTAuthService()

	// repository
	todoRepository := repository.NewTodoRepository()
	userRepository := repository.NewUserRepository()

	// user
	userService := service.NewUserService(userRepository, todoRepository, db, validator)
	userController := controller.NewUserController(userService, jwtService)

	// todo
	todoService := service.NewTodoService(todoRepository, userRepository, db, validator)
	todoController := controller.NewTodoController(todoService)

	router := app.NewRouter(todoController, userController)
	return router
}

func All() (*gin.Engine, *sql.DB) {
	db := DB()
	TruncateAll(db)
	router := Router(db)

	return router, db
}
