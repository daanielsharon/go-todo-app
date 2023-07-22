package middleware

import "github.com/gin-gonic/gin"

func SetResponseHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
	}
}
