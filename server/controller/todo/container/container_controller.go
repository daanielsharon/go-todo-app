package containercontr

import "github.com/gin-gonic/gin"

type ContainerController interface {
	UpdatePriority(ctx *gin.Context)
}
