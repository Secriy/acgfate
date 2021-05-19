package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"passwd"`
	DB       int    `mapstructure:"db"`
}

// RDB Redis缓存客户端单例
var RDB *redis.Client

// InitRedisClient 在中间件中初始化redis链接
func InitRedisClient(conf RedisConf) {
	// conf := config.Conf.Redis // 配置
	RDB = redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password,
		DB:       conf.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()

	if err != nil {
		fmt.Println("连接Redis不成功", err)
	}
}
