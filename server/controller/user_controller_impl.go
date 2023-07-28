package controller

import (
	"net/http"
	"server/exception"
	"server/helper"
	"server/model/web"
	"server/service"
	"server/util"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService    service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService:    userService,
	}
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	var req web.UserCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	res := c.UserService.Create(ctx, &req)
	helper.WriteToResponseBody(ctx, res)
}

func (c *UserControllerImpl) Login(ctx *gin.Context) {
	var req web.UserGetRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	user := c.UserService.Get(ctx, &req)

	res := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}

	token := util.NewJWTAuth().TokenGenerate(req.Username)
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
