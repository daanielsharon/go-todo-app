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

	// router
	router := app.NewRouter(todoController, userController)
	router.Run(":8080")
}
