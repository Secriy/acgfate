package config

import (
	"fmt"

	"acgfate/model"
	"github.com/spf13/viper"
)

type config struct {
	DSN   string `mapstructure:"dsn"`
	Mode  string `mapstructure:"mode"`
	Redis struct {
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"passwd"`
		DB       string `mapstructure:"db"`
	} `mapstructure:"redis"`
	JWT struct {
		Secret         string `mapstructure:"secret"`
		ExpireDuration int    `mapstructure:"expire_duration"`
	} `mapstructure:"jwt"`
	Mail struct {
		Smtp     string `mapstructure:"smtp"`
		Port     int    `mapstructure:"port"`
		Sender   string `mapstructure:"sender"`
		Password string `mapstructure:"passwd"`
	} `maostructure:"mail"`
}

var Conf config

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
	// Initialize database
	model.InitDatabase(Conf.DSN)
}
