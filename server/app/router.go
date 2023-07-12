package app

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func NewRouter(todo controller.TodoController, user controller.UserController) *gin.Engine {
	r := setupRouter()

	// user
	r.POST("/users/register", user.Register)
	r.POST("/users/login", user.Login)

	// todo
	r.GET("/todo/:username", todo.GetTodoByUsername)
	r.POST("/todo/create-todo", todo.CreateTodo)
	r.DELETE("/todo", todo.RemoveTodo)

	return r
}
