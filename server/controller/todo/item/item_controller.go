package itemcontr

import "github.com/gin-gonic/gin"

type ItemController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetByUsername(ctx *gin.Context)
	Remove(ctx *gin.Context)
}
