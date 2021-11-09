package cache

import (
	"context"
	"fmt"
	"time"

	"acgfate/config"
	"github.com/go-redis/redis/v8"
)

// RDB Redis缓存客户端单例
var rdb *redis.Client

// InitRedis 在中间件中初始化redis链接
func InitRedis(conf *config.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Passwd,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return
}

func CloseRedis() {
	_ = rdb.Close()
}
