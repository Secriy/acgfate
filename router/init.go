package router

import (
	"acgfate/config"
	_ "acgfate/docs"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize the router
func InitRouter() *gin.Engine {
	conf := config.Conf
	// Router Mode
	gin.SetMode(conf.Mode)
	// New Router
	r := gin.Default()
	// Middleware
	r.Use(middleware.Session(conf.Session.Secret, conf.RedisConf.Host))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	r.Use(gin.Recovery())
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
