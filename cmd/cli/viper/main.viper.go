package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
	Security struct {
		Jwt struct {
			Secret string `mapstructure:"secret"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration file: %w", err))
	}

	fmt.Println("Server running is port:", viper.GetInt("server.port"))
	fmt.Println("JWT running is secret:", viper.GetString("security.jwt.secret"))

	// Unmarshal config
	var config Config
	if err := viper.Unmarshal(&config); err != nil { // tro toi dia chi config
		fmt.Printf("Unable to decode configuration %s \n", err)
	}

	fmt.Println("Config Port:: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database user: %s host: %s password: %s dbname: %s \n", db.User, db.Host, db.Password, db.DbName)
	}

	fmt.Println("JWT key:: ", config.Security.Jwt.Secret)
}
