package main

import (
	"server/app"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDatabase("todoapp")
	validator := validator.New()
	timeout := time.Duration(1) * time.Second

	userController, containerController, itemController := app.NewAppSetup(db, validator, timeout)

	// router
	router := app.NewRouter(containerController, itemController, userController)
	router.Run(":8080")
}
