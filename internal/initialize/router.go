package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/global"
	"github.com/ntquang/ecommerce/internal/routers"
	"net/http"
)

func Initrouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Logger(), gin.Recovery())
	}

	//middlewares
	//r.Use() //logging
	//r.Use() // cross
	//r.Use() // limiter in global

	//r.Use(middlewares.NewRateLimiter().GlobalRateLimiter())
	//r.GET("/ping/100", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"message": "pong 100",
	//	})
	//})
	//
	//r.Use(middlewares.NewRateLimiter().UserPrivateAPIRateLimiter())
	//r.GET("/ping/50", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"message": "pong 50",
	//	})
	//})
	//
	//r.Use(middlewares.NewRateLimiter().PublicAPIRateLimiter())
	//r.GET("/ping/80", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"message": "pong 80",
	//	})
	//})

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User
	oauth2Router := routers.RouterGroupApp.Oauth2
	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus") //tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}
	{
		oauth2Router.InitOauth2Router(MainGroup)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route Not Found"})
	})

	return r
}
