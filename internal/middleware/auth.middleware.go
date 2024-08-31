package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/response"
)

func Authentication() gin.HandlerFunc {
	// Add authentication logic here
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			response.ErrorResponse(c, response.ErrTokenHeadersInvalid, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
