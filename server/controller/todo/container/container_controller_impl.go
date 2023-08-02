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

func (c *ContainerControllerImpl) UpdatePriority(ctx *gin.Context) {
	var req web.TodoUpdatePriority

	updateId := ctx.Param("groupId")
	id, err := strconv.Atoi(updateId)
	req.OriginID = int64(id)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.ContainerService.UpdatePriority(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}
