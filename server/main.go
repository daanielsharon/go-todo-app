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

	// repository
	todoRepository := repository.NewTodoRepository()
	userRepository := repository.NewUserRepository()

	// user
	userService := service.NewUserService(userRepository, todoRepository, db)
	userController := controller.NewUserController(userService)

	// todo
	todoService := service.NewTodoService(todoRepository, userRepository, db, validator)
	todoController := controller.NewTodoController(todoService)

	// router
	router := app.NewRouter(todoController, userController)
	router.Run(":8080")
}
