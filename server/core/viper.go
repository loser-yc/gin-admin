package core

import (
	"github.com/spf13/viper"
	"go-admin/config"
	"log"
)

var Conf config.Config

func SetConfig()  {
	//设置viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	//获取配置
	viper.Unmarshal(&Conf)
}
