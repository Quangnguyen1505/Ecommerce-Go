package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/internal/controller/account"
	"github.com/ntquang/ecommerce/internal/middleware"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	//non-dependency Injection
	// userRepository := repo.NewUserRepository()
	// userService := services.NewUserService(userRepository)
	// userHandleNondenpendency := controller.NewUserController(userService)

	//WIRE 	go
	//Denpendency Injection
	// userController, _ := wire.InitUserRouterHanlder()
	userRouterPublic := Router.Group("/user")
	{
		// userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/verifyOTP", account.Login.VerifyOTP)
		userRouterPublic.POST("/login", account.Login.Login)
		userRouterPublic.POST("/updatePass", account.Login.UpdatePasswordRegister)
	}

	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middleware.Authentication())
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/info")
		userRouterPrivate.POST("/two-factor/setup", account.User2fa.SetupTwoFactorAuth)
	}
}
