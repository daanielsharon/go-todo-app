package app

import (
	"database/sql"
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

func NewAppSetup(db *sql.DB, validator *validator.Validate, timeout time.Duration) (controller.UserController, containercontr.ContainerController, itemcontr.ItemController) {
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

	return userController, containerController, itemController
}
