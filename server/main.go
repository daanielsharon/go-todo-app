package main

import (
	"server/app"
	"server/controller"
	containercontr "server/controller/todo/container"
	itemcontr "server/controller/todo/item"
	"server/repository"
	containerrepo "server/repository/todo/container"
	itemrepo "server/repository/todo/item"
	"server/service"
	containerserv "server/service/todo/container"
	itemserv "server/service/todo/item"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDatabase()
	validator := validator.New()
	timeout := time.Duration(1) * time.Second

	// repository
	itemRepository := itemrepo.NewItemRepository()
	containerRepository := containerrepo.NewContainerRepository()
	userRepository := repository.NewUserRepository()

	// user
	userService := service.NewUserService(userRepository, containerRepository, db, validator)
	userController := controller.NewUserController(userService)

	// todoContainer
	containerService := containerserv.NewContainerService(containerRepository, db, validator, timeout)
	containerController := containercontr.NewContainerController(containerService)

	// todoItem
	itemService := itemserv.NewItemService(itemRepository, containerRepository, userRepository, db, validator, timeout)
	itemController := itemcontr.NewItemController(itemService)

	// router
	router := app.NewRouter(containerController, itemController, userController)
	router.Run(":8080")
}
