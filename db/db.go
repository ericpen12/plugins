package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
)

var Client *gorm.DB

type config struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
	sslMode  string
}

func Connect() error {
	var err error
	c := loadConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.user, c.password, c.host+":"+c.port, c.dbname)
	Client, err = gorm.Open("mysql", dsn)
	return err
}

func loadConfig() *config {
	return &config{
		host:     viper.GetString("db.host"),
		port:     viper.GetString("db.port"),
		user:     viper.GetString("db.user"),
		dbname:   viper.GetString("db.dbname"),
		password: viper.GetString("db.password"),
		sslMode:  viper.GetString("db.sslMode"),
	}
}
