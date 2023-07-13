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

	api := r.Group("api")
	{
		v1 := api.Group("v1")
		{
			// user
			users := v1.Group("/users")
			{
				users.POST("/register", user.Register)
				users.POST("/login", user.Login)
			}

			// todo
			todos := v1.Group("/todo")
			{
				todos.GET("/todo/:username", todo.GetTodoByUsername)
				todos.POST("/todo/create-todo", todo.CreateTodo)
				todos.DELETE("/todo", todo.RemoveTodo)
			}
		}
	}

	return r
}
