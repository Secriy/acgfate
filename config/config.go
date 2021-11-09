package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name          string `mapstructure:"name"`
	Version       string `mapstructure:"version"`
	Mode          string `mapstructure:"mode"`
	Port          int    `mapstructure:"port"`
	SessionSecret string `mapstructure:"session_secret"`
	StartTime     string `mapstructure:"start_time"`
	MachineID     int64  `mapstructure:"machine_id"`
	*LogConfig    `mapstructure:"log"`
	*MySQLConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
	*EmailConfig  `mapstructure:"email"`
	*JWTConfig    `mapstructure:"jwt"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Passwd       string `mapstructure:"passwd"`
	DB           string `mapstructure:"db"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Passwd       string `mapstructure:"passwd"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type EmailConfig struct {
	Smtp   string `mapstructure:"smtp"`
	Port   int    `mapstructure:"port"`
	Passwd string `mapstructure:"passwd"`
	Sender string `mapstructure:"sender"`
}

type JWTConfig struct {
	Secret         string `mapstructure:"secret"`
	ExpireDuration string `mapstructure:"expire_duration"`
}

// Init is used to read configuration file.
func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error while reading configration file: %s \n", err)
		return err
	}

	if err := viper.Unmarshal(Conf); err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("configuration was changed...")
		if err := viper.Unmarshal(Conf); err != nil {
			return
		}
	})
	return nil
}
