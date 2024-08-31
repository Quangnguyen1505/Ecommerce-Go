package initialize

import (
	"fmt"

	"github.com/ntquang/ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	InitConfig()
	fmt.Println("Security token is::", global.Config.Security.Jwt.Secret)
	InitLogger()
	global.Logger.Error("Error log", zap.Int("line", 1))
	InitPostgresql()
	InitRedis()
	Initrouter()
	r := Initrouter()
	r.Run(":8081")
}
