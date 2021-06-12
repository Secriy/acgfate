package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"passwd"`
	DB       int    `mapstructure:"db"`
}

type JwtConf struct {
	Secret         string `mapstructure:"secret"`
	ExpireDuration int    `mapstructure:"expire_duration"`
}

type MailConf struct {
	Smtp     string `mapstructure:"smtp"`
	Port     int    `mapstructure:"port"`
	Sender   string `mapstructure:"sender"`
	Password string `mapstructure:"passwd"`
}

type Config struct {
	DSN   string    `mapstructure:"dsn"`
	Mode  string    `mapstructure:"mode"`
	Redis RedisConf `mapstructure:"redis"`
	JWT   JwtConf   `mapstructure:"jwt"`
	Email MailConf  `maostructure:"email"`
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
