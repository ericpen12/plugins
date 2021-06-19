package settings

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var (
	httpPort = 9090
	appName  = ""
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

var (
	configName = "config"
	configType = "yaml"
	configPath = "./"
)

func SetConfigPath(path string) {
	list := strings.Split(filepath.Base(path), ".")
	configPath = filepath.Dir(path)
	configName = list[0]
	if len(list) > 1 {
		configType = list[1]
	}
}

func loadConfigFile() error {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath("./config/")
	return viper.ReadInConfig()
}

func HttpPort() int {
	return httpPort
}

func AppName() string {
	return appName
}
