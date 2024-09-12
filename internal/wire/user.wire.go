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
		repo.NewUserAuthRepository,
		services.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
