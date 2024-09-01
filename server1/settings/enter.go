package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"server1/output"
)

var Config = &ConfigClass{}

func init() {
	initViper()
	initPort()
	initRoutes()
}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(output.Error() + "Failed to read config file")
	} else {
		fmt.Println(output.Info() + "Config file read successfully")
	}
}

func initPort() {
	Config.Port = viper.GetInt64("port")
}
