package router

import (
	v1 "acgfate/api/http/v1"
	"github.com/gin-gonic/gin"
)

// InitCheckRouter 初始化校验相关路由
func InitCheckRouter(r *gin.RouterGroup) {
	// 公开路由组
	pub := r.Group("check")
	{
		pub.GET("username", v1.CheckUsername) // 检查用户名是否存在
		pub.GET("email", v1.CheckEmail)       // 检查邮箱是否存在
	}
}
