package setup

import (
	"database/sql"
	"server/app"
	"server/controller"
	containercontr "server/controller/todo/container"
	itemcontr "server/controller/todo/item"
	"server/helper"
	"server/repository"
	containerrepo "server/repository/todo/container"
	itemrepo "server/repository/todo/item"
	"server/service"
	containerserv "server/service/todo/container"
	itemserv "server/service/todo/item"
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

	router := app.NewRouter(containerController, itemController, userController)
	return router
}

func All() (*gin.Engine, *sql.DB) {
	db := DB()
	TruncateAll(db)
	router := Router(db)

	return router, db
}
