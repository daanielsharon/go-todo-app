package middleware

import (
	"server/exception"
	"server/util"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("token")
		if err != nil {
			panic(exception.NewUnauthorizedError("token is not provided!"))
		}

		token, err := util.NewToken().TokenValidate(authCookie)
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
