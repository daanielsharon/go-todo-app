package exception

import (
	"net/http"
	"server/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	if unauthorizedErrors(c, err) {
		c.Abort()
		return
	}

	if validationErrors(c, err) {
		c.Abort()
		return
	}

	if notFoundErrors(c, err) {
		c.Abort()
		return
	}

	if serviceValidationErrors(c, err) {
		c.Abort()
		return
	}

	internalServerErrors(c, err)
	return
}

func unauthorizedErrors(c *gin.Context, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		res := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   exception.Error,
		}

		c.JSON(http.StatusUnauthorized, res)
		return true
	} else {
		return false
	}
}

func validationErrors(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		res := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		c.JSON(http.StatusBadRequest, res)
		return true

	} else {
		return false
	}
}

func serviceValidationErrors(c *gin.Context, err interface{}) bool {
	exception, ok := err.(ValidationError)

	if ok {
		res := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error,
		}

		c.JSON(http.StatusBadRequest, res)
		return true
	} else {
		return false
	}
}

func notFoundErrors(c *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		res := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		c.JSON(http.StatusNotFound, res)
		return true
	} else {
		return false
	}
}

func internalServerErrors(c *gin.Context, err interface{}) {
	res := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	c.JSON(http.StatusInternalServerError, res)
}
