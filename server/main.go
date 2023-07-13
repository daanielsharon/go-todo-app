package main

import (
	"server/app"
	"server/controller"
	"server/repository"
	"server/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDatabase()
	validator := validator.New()

	// user
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	// todo
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, userRepository, db, validator)
	todoController := controller.NewTodoController(todoService)

	// router
	router := app.NewRouter(todoController, userController)
	router.Run(":8080")
}
