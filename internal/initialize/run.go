package initialize

import (
	"fmt"

	"github.com/ntquang/ecommerce/global"
)

func Run() {
	InitConfig()
	InitLogger()
	InitPostgresql()
	InitRedis()
	Initrouter()
	r := Initrouter()
	port := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(port)
}
