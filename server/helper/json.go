package helper

import (
	"net/http"
	"server/model/web"

	"github.com/gin-gonic/gin"
)

func WriteToResponseBody(c *gin.Context, response interface{}) {
	res := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	c.JSON(http.StatusOK, res)
}
