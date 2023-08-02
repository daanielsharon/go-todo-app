package itemcontr

import (
	"net/http"
	"server/exception"
	"server/helper"
	"server/model/web"
	itemserv "server/service/todo/item"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemControllerImpl struct {
	Service itemserv.ItemService
}

func NewItemController(service itemserv.ItemService) ItemController {
	return &ItemControllerImpl{
		Service: service,
	}
}

func (c *ItemControllerImpl) Create(ctx *gin.Context) {
	var req web.TodoCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.Service.Create(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *ItemControllerImpl) Update(ctx *gin.Context) {
	var req web.TodoUpdateRequest

	updateId := ctx.Param("todoId")
	id, err := strconv.Atoi(updateId)
	req.ID = int64(id)

	err = ctx.ShouldBindJSON(&req)

	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.Service.Update(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *ItemControllerImpl) GetByUsername(ctx *gin.Context) {
	var req web.TodoGetRequest

	username := ctx.Param("username")
	req.Username = username

	res := c.Service.GetByUsername(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *ItemControllerImpl) Remove(ctx *gin.Context) {
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
