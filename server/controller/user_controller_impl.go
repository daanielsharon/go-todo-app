package controller

import (
	"net/http"
	"server/exception"
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
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.Service.CreateUsername(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *UserControllerImpl) Login(ctx *gin.Context) {
	var req web.UserGetUsernameRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	user := c.Service.GetUsername(ctx, &req)

	res := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	token := c.JWTService.TokenGenerate(req.Username)
	ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, res)
}

func (c *UserControllerImpl) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "", "", false, true)

	res := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	ctx.JSON(http.StatusOK, res)
}
