package app

import (
	"server/controller"
	"server/exception"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func NewRouter(todo controller.TodoController, user controller.UserController) *gin.Engine {
	r := setupRouter()
	r.Use(gin.CustomRecovery(exception.ErrorHandler))

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
			todos := v1.Group("/todo", middleware.Authorize())
			{
				todos.GET("/:username", todo.GetTodoByUsername)
				todos.POST("/", todo.CreateTodo)
				todos.DELETE("/:id", todo.RemoveTodo)
			}
		}
	}

	return r
}
