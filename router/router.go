package router

import (
	config "acgfate/config"
	_ "acgfate/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	// Router Mode
	gin.SetMode(config.Conf.Mode)

	// New Router
	r := gin.Default()

	// Middleware
	r.Use(gin.Recovery())

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Router Group
	rGroup := r.Group("/api/v1")

	InitUserRouter(rGroup)
	InitWordsRouter(rGroup)

	return r
}
