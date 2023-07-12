package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteToResponseBody(c *gin.Context, res interface{}) {
	c.JSON(http.StatusOK, res)
}
