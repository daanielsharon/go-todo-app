package middleware

import (
	"server/exception"
	"server/service"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("token")
		if err != nil {
			panic(exception.NewUnauthorizedError("token is not provided!"))
		}

		token, err := service.NewJWTAuthService().TokenValidate(authCookie)
		if err != nil {
			panic(exception.NewUnauthorizedError(err.Error()))
		}

		if token.Valid {
			c.Next()
		} else {
			panic(exception.NewUnauthorizedError("invalid token"))
		}
	}
}
