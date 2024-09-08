//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/ntquang/ecommerce/internal/controller"
	"github.com/ntquang/ecommerce/internal/repo"
	"github.com/ntquang/ecommerce/internal/services"
)

func InitUserRouterHanlder() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		services.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
