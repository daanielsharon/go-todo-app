package controller

import (
	"net/http"
	"server/exception"
	"server/helper"
	"server/model/web"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoControllerImpl struct {
	Service service.TodoService
}

func NewTodoController(service service.TodoService) TodoController {
	return &TodoControllerImpl{
		Service: service,
	}
}

func (c *TodoControllerImpl) CreateTodo(ctx *gin.Context) {
	var req web.TodoCreateRequest

	err := ctx.ShouldBindJSON(&req)
	helper.PanicIfError(err)

	res := c.Service.CreateTodo(ctx, &req)

	helper.WriteToResponseBody(ctx, res)
}

func (c *TodoControllerImpl) GetTodoByUsername(ctx *gin.Context) {
	var req web.TodoGetRequest

	username := ctx.Param("username")
	req.Username = username

	res := c.Service.GetTodoByUsername(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *TodoControllerImpl) RemoveTodo(ctx *gin.Context) {
	var req web.TodoDeleteRequest

	deleteId := ctx.Param("id")
	id, err := strconv.Atoi(deleteId)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}
	req.ID = int64(id)

	c.Service.RemoveTodo(ctx, &req)

	res := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	ctx.JSON(http.StatusOK, res)
}
