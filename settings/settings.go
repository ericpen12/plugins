package settings

import "github.com/spf13/viper"

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config/")
	return viper.ReadInConfig()
}
