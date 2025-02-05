package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	InitConfig()
	InitLogger()
	// InitPostgresql()
	InitPostgresqlC()
	InitRedis()
	InitServiceInterface()
	InitKafa()
	InitOauth2()
	// Initrouter()
	r := Initrouter()
	// port := fmt.Sprintf(":%v", global.Config.Server.Port)
	// r.Run(port)
	return r
}
