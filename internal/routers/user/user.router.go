package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ntquang/ecommerce/internal/wire"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	//non-dependency Injection
	// userRepository := repo.NewUserRepository()
	// userService := services.NewUserService(userRepository)
	// userHandleNondenpendency := controller.NewUserController(userService)

	//WIRE 	go
	//Denpendency Injection
	userController, _ := wire.InitUserRouterHanlder()
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/regiser", userController.Register)
	}

	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/info")
	}
}
