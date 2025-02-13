package routers

import (
	"github.com/ntquang/ecommerce/internal/routers/manage"
	"github.com/ntquang/ecommerce/internal/routers/user"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	User   user.UserRouterGroup
}

var RouterGroupApp = &RouterGroup{}
