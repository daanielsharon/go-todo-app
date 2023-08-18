package containercontr

import (
	"server/exception"
	"server/helper"
	"server/model/web"
	containerserv "server/service/todo/container"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContainerControllerImpl struct {
	ContainerService containerserv.ContainerService
}

func NewContainerController(containerService containerserv.ContainerService) ContainerController {
	return &ContainerControllerImpl{
		ContainerService: containerService,
	}
}

func (c *ContainerControllerImpl) Create(ctx *gin.Context) {
	var req web.ContainerCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.ContainerService.Create(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *ContainerControllerImpl) UpdatePriority(ctx *gin.Context) {
	var req web.TodoUpdatePriority

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	updateId := ctx.Param("groupId")
	id, err := strconv.Atoi(updateId)
	req.OriginID = int64(id)

	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.ContainerService.UpdatePriority(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *ContainerControllerImpl) Delete(ctx *gin.Context) {
	var req web.ContainerDeleteRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	deleteId := ctx.Param("groupId")
	id, err := strconv.Atoi(deleteId)

	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	req.ID = int64(id)

	c.ContainerService.Delete(ctx, &req)
}
