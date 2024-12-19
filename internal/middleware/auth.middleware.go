package middleware

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/internal/utlis/auth"
)

func Authentication() gin.HandlerFunc {
	// Add authentication logic here
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		log.Println("Url request ::", url)

		//get token in header
		JwtToken, ok := auth.ExtractBearerToken(c)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"code":        40001,
				"err":         "Unauthorized",
				"description": "",
			})
			return
		}

		//validate jwt token by subject
		claims, err := auth.VerifyToken(JwtToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"code":        40001,
				"err":         "Invalid Token",
				"description": "" + err.Error(),
			})
			return
		}

		log.Println("Claims UUID:: ", claims.Subject) //11clitoken...
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
