package setup

import (
	"database/sql"
	"server/app"
	"server/controller"
	"server/helper"
	"server/repository"
	containerTodo "server/repository/todo/container"
	itemTodo "server/repository/todo/item"
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

	// repository
	itemRepository := itemTodo.NewItemRepository()
	containerRepository := containerTodo.NewContainerRepository()
	userRepository := repository.NewUserRepository()

	// user
	userService := service.NewUserService(userRepository, containerRepository, db, validator)
	userController := controller.NewUserController(userService)

	// todo
	todoService := service.NewTodoService(itemRepository, containerRepository, userRepository, db, validator)
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
