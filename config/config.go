package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Version string `mapstructure:"version"`
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	DSN     string `mapstructure:"dsn"`
	// Redis
	RedisConf struct {
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"passwd"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
	// Auth
	Session struct {
		Secret string `mapstructure:"secret"`
	}
	JWTConf struct {
		Secret         string `mapstructure:"secret"`
		ExpireDuration int    `mapstructure:"expire_duration"`
	} `mapstructure:"jwt"`
	// Email
	EmailConf struct {
		Smtp     string `mapstructure:"smtp"`
		Port     int    `mapstructure:"port"`
		Sender   string `mapstructure:"sender"`
		Password string `mapstructure:"passwd"`
	} `mapstructure:"email"`
}

var Conf Config

// ReadConfig is used to read configuration file
func ReadConfig() {
	// Read config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// 读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error read configration file: %s \n", err))
	}
	// 反序列化
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarshal configration file: %s \n", err))
	}
}
