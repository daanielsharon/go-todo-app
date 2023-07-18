package exception

import (
	"net/http"
	"server/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	if validationErrors(c, err) {
		return
	}

	if notFoundError(c, err) {
		return
	}

	if serviceValidationErrors(c, err) {
		return
	}

	internalServerError(c, err)
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
	exception, ok := err.(ServiceValidationError)

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

func notFoundError(c *gin.Context, err interface{}) bool {
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

func internalServerError(c *gin.Context, err interface{}) {
	res := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	c.JSON(http.StatusInternalServerError, res)
}
