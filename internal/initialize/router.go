package initialize

import (
	"github.com/gin-gonic/gin"
	c "github.com/ntquang/ecommerce/internal/controller"
	"github.com/ntquang/ecommerce/internal/middleware"
)

func Initrouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Authentication())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUser)
		// v1.POST("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.PUT("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	return r
}
