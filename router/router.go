package router

import (
	config "acgfate/conf"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(gin.Recovery())

	// Router Mode
	gin.SetMode(config.Conf.Mode)

	// Router Group
	rGroup := r.Group("/api/v1")

	InitUserRouter(rGroup)

	return r
}
