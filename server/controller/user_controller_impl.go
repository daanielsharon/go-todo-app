package controller

import (
	"net/http"
	"server/helper"
	"server/model/web"
	"server/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return &UserControllerImpl{
		service: s,
	}
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	var req web.UserCreateUsernameRequest

	username := ctx.Param("username")
	req.Username = username

	res := c.service.CreateUsername(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *UserControllerImpl) Login(ctx *gin.Context) {
	var req web.UserGetUsernameRequest

	username := ctx.Param("username")
	req.Username = username

	c.service.GetUsername(ctx, &req)

	res := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	ctx.JSON(http.StatusOK, res)
}
