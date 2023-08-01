package main

import (
	"server/app"
	"server/controller"
	"server/repository"
	containerTodo "server/repository/todo/container"
	itemTodo "server/repository/todo/item"
	"server/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDatabase()
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

	// router
	router := app.NewRouter(todoController, userController)
	router.Run(":8080")
}
