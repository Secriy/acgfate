package router

import (
	"fmt"

	"acgfate/config"
	_ "acgfate/docs"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Init initialize the router
func Init(conf *config.RedisConfig, secret string) *gin.Engine {
	r := gin.New()

	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.Use(middleware.Session(secret, fmt.Sprintf("%s:%d",
		conf.Host,
		conf.Port,
	)))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Router Group
	v1 := r.Group("/api/v1")

	InitUserRouter(v1) // 初始化UserInfo路由
	// InitTaskRouter(v1)  // 初始化任务相关路由
	// InitCheckRouter(v1) // 初始化校验相关路由
	// InitWordsRouter(v1) // 初始化Words路由

	return r
}
