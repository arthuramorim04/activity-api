package middlewares

import (
	"github.com/arthuramorim04/go-books-api.git/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schemas = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
			return
		}

		token := header[len(Bearer_schemas):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
			return
		}
	}
}
