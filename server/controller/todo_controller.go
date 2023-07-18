package controller

import "github.com/gin-gonic/gin"

type TodoController interface {
	CreateTodo(ctx *gin.Context)
	GetTodoByUsername(ctx *gin.Context)
	RemoveTodo(ctx *gin.Context)
}
