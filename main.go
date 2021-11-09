package main

import (
	"fmt"

	"acgfate/cache"
	"acgfate/config"
	"acgfate/database"
	"acgfate/router"
	"acgfate/utils/logger"
	"acgfate/utils/snowflake"
	"github.com/gin-gonic/gin"
)

// @title ACG.Fate API
// @version 0.1.1
// @description The RESTFul API of Server
// @host 127.0.0.1:3000
// @BasePath /api/v1
func main() {
	if err := config.Init(); err != nil {
		fmt.Printf("Init the configuration failed: %e\n", err)
		return
	}
	appConf := config.Conf

	if err := logger.Init(appConf.LogConfig, appConf.Mode); err != nil {
		fmt.Printf("Init the logger failed: %e\n", err)
		return
	}
	if err := database.InitMySQL(appConf.MySQLConfig); err != nil {
		fmt.Printf("Init the MySQL connection failed: %e\n", err)
		return
	}
	defer database.CloseMySQL() // close MySQL connection
	if err := cache.InitRedis(appConf.RedisConfig); err != nil {
		fmt.Printf("Init the Redis connection failed: %e\n", err)
		return
	}
	defer cache.CloseRedis() // close Redis connection

	// 初始化雪花算法结点
	if err := snowflake.Init(appConf.StartTime, appConf.MachineID); err != nil {
		fmt.Printf("Init the Snowflake node failed: %e\n", err)
		return
	}

	gin.SetMode(appConf.Mode)

	r := router.Init(appConf.RedisConfig, appConf.SessionSecret)

	err := r.Run(fmt.Sprintf(":%d", appConf.Port))
	if err != nil {
		fmt.Printf("Error run Gin router: %e \n", err)
		return
	}
}
