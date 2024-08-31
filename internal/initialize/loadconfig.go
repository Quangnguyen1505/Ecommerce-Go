package initialize

import (
	"fmt"

	"github.com/ntquang/ecommerce/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration file: %w", err))
	}

	// fmt.Println("Server running is port:", viper.GetInt("server.port"))
	// fmt.Println("JWT running is secret:", viper.GetString("security.jwt.secret"))
	// Unmarshal config
	// var config Config
	if err := viper.Unmarshal(&global.Config); err != nil { // tro toi dia chi config
		fmt.Printf("Unable to decode configuration %s \n", err)
	}
}
