package controller

import (
	"net/http"
	"server/helper"
	"server/model/web"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoControllerImpl struct {
	service service.TodoService
}

func NewTodoController(s service.TodoService) TodoController {
	return &TodoControllerImpl{
		service: s,
	}
}

func (c *TodoControllerImpl) CreateTodo(ctx *gin.Context) {
	var req web.TodoCreateRequest

	err := ctx.ShouldBindJSON(&req)
	helper.PanicIfError(err)

	res := c.service.CreateTodo(ctx, &req)

	helper.WriteToResponseBody(ctx, res)
}

func (c *TodoControllerImpl) GetTodoByUsername(ctx *gin.Context) {
	var req web.TodoGetRequest

	username := ctx.Param("username")
	req.Username = username

	res := c.service.GetTodoByUsername(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *TodoControllerImpl) RemoveTodo(ctx *gin.Context) {
	var req web.TodoDeleteRequest

	deleteId := ctx.Param("id")
	id, err := strconv.Atoi(deleteId)
	helper.PanicIfError(err)
	req.ID = int64(id)

	c.service.RemoveTodo(ctx, &req)

	res := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	ctx.JSON(http.StatusOK, res)
}
