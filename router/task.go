package router

import (
	v1 "acgfate/api/http/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

// InitTaskRouter 初始化任务相关路由
func InitTaskRouter(r *gin.RouterGroup) {
	gp := r.Group("task")
	gp.Use(middleware.JWTAuthRequired(), middleware.IsVerified(), middleware.IsBanned())
	{
		gp.GET("sign", v1.TaskSign)
	}
}
