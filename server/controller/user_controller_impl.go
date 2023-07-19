package controller

import (
	"net/http"
	"server/helper"
	"server/model/web"
	"server/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	Service    service.UserService
	JWTService service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &UserControllerImpl{
		Service:    userService,
		JWTService: jwtService,
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

	token := c.JWTService.TokenGenerate(req.Username)
	ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, res)
}
