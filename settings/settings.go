package settings

import (
	"github.com/spf13/viper"
)

var (
	httpPort = 9090
	appName = ""
)

func Init() error {
	err := loadConfigFile()
	if err != nil {
		return err
	}
	if viper.GetInt("app.port") != 0 {
		httpPort = viper.GetInt("app.port")
	}
	appName = viper.GetString("app.name")
	return nil
}

func loadConfigFile() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config/")
	return viper.ReadInConfig()
}

func HttpPort() int {
	return httpPort
}

func AppName() string {
	return appName
}