package router

import (
	"fmt"

	apiv1 "acgfate/api/http/v1"
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

	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", apiv1.UserRegister)   // 注册
		v1.POST("user/login", apiv1.UserLogin)         // 登录
		v1.GET("category/:name", apiv1.CategoryDetail) // 分区信息
		v1.GET("category/list", apiv1.CategoryList)    // 分区列表
		v1.GET("word/list", apiv1.WordList)            // 文字列表
		v1.GET("word/:id", apiv1.WordDetail)           // 文字详情

		// 需要鉴权的路由组
		v1.Use(middleware.AuthRequired())

		v1.GET("user/info", apiv1.UserInfo)            // 获取个人信息
		v1.POST("word/post", apiv1.WordPost)           // 发表文字
		v1.DELETE("word/:id/delete", apiv1.WordDelete) // 发表文字
	}
	return r
}
