package app

import (
	"net/http"
	"server/controller"
	"server/exception"
	"server/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.CustomRecovery(exception.ErrorHandler))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowHeaders:     []string{"Content-Type", "Accept"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
		AllowWebSockets:  false,
		MaxAge:           12 * time.Hour,
	}))

	return r
}

func NewRouter(todo controller.TodoController, user controller.UserController) *gin.Engine {
	r := setupRouter()

	api := r.Group("api", middleware.SetResponseHeader())
	{
		v1 := api.Group("v1")
		{
			// user
			users := v1.Group("/users")
			{
				users.POST("/register", user.Register)
				users.POST("/login", user.Login)
				users.POST("/logout", middleware.Authorize(), user.Logout)
			}

			// todo
			todos := v1.Group("/todo", middleware.Authorize())
			{
				todos.GET("/:username", todo.GetByUsername)
				todos.POST("/", todo.Create)
				todos.PATCH("/", todo.Update)
				todos.DELETE("/:todoId", todo.Remove)
			}
		}
	}

	return r
}
