package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.GET("/login")
	}

	adminRouterPrivate := Router.Group("/admin")
	// adminRouterPrivate.Use(Limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.GET("/active_user")
	}
}
