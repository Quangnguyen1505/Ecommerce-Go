package initialize

import (
	"github.com/ntquang/ecommerce/global"
	"github.com/ntquang/ecommerce/internal/database"
	"github.com/ntquang/ecommerce/internal/services"
	"github.com/ntquang/ecommerce/internal/services/imple"
)

func InitServiceInterface() {
	queries := database.New(global.Pdbc)
	// Init Interface here

	//user service interface
	services.InitUserLogin(imple.NewUserLoginImpl(queries))
	//user admin interface
	services.InitUserAdmin(imple.NewUserAdminImpl(queries))
	//..

}
