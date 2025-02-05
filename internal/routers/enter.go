package routers

import (
	"github.com/ntquang/ecommerce/internal/routers/manage"
	"github.com/ntquang/ecommerce/internal/routers/oauth2"
	"github.com/ntquang/ecommerce/internal/routers/user"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	User   user.UserRouterGroup
	Oauth2 oauth2.Oauth2RouterGroup
}

var RouterGroupApp = &RouterGroup{}
