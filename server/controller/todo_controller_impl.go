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

func (c *TodoControllerImpl) Create(ctx *gin.Context) {
	var req web.TodoCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.Service.Create(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *TodoControllerImpl) Update(ctx *gin.Context) {
	var req web.TodoUpdateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.Service.Update(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *TodoControllerImpl) GetByUsername(ctx *gin.Context) {
	var req web.TodoGetRequest

	username := ctx.Param("username")
	req.Username = username

	res := c.Service.GetByUsername(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *TodoControllerImpl) Remove(ctx *gin.Context) {
	var req web.TodoDeleteRequest

	deleteId := ctx.Param("todoId")
	id, err := strconv.Atoi(deleteId)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}
	req.ID = int64(id)

	c.Service.Remove(ctx, &req)

	res := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	ctx.JSON(http.StatusOK, res)
}
