package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"testing"
)

func initSettings() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	return viper.ReadInConfig()
}

func TestInitLogWithYaml(t *testing.T) {
	err := initSettings()
	if err != nil {
		t.Errorf("cannot initialize setting, %v", err)
		return
	}
	Init()
	zap.L().Info("init log with a yaml file")
}

func TestInitLogWithoutYaml(t *testing.T) {
	Init()
	zap.L().Info("init log with a yaml file")
}

func Test_logPath(t *testing.T) {
	level := "error"
	ret := logPathWithAppNameAndLevel(level)
	t.Log(ret)
}
