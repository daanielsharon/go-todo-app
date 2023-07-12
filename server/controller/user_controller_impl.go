package controller

import (
	"net/http"
	"server/helper"
	"server/model/web"
	"server/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	var req web.UserCreateUsernameRequest

	err := ctx.ShouldBindJSON(&req)
	helper.PanicIfError(err)

	res := c.Service.CreateUsername(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *UserControllerImpl) Login(ctx *gin.Context) {
	var req web.UserGetUsernameRequest

	err := ctx.ShouldBindJSON(&req)
	helper.PanicIfError(err)

	c.Service.GetUsername(ctx, &req)

	res := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	ctx.JSON(http.StatusOK, res)
}
