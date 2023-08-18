package containercontr

import "github.com/gin-gonic/gin"

type ContainerController interface {
	Create(ctx *gin.Context)
	UpdatePriority(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
