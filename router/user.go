package router

import (
	v12 "acgfate/api/http/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	// 公开路由组
	pubGroup := r.Group("user")
	{
		pubGroup.POST("register", v12.UserRegister) // 注册
		pubGroup.POST("login", v12.UserLogin)       // 登录
	}

	// 需要鉴权的路由组
	authGroup := pubGroup.Use(middleware.JWTAuthRequired())
	{
		authGroup.GET("me", v12.UserMe)         // 获取个人信息
		authGroup.PUT("update", v12.UserUpdate) // 更新个人信息
	}

	// 邮箱路由组
	mailGroup := r.Group("mail").Use(middleware.JWTAuthRequired())
	{
		mailGroup.GET("verify", v12.MailSend)    // 发送验证码
		mailGroup.POST("verify", v12.MailVerify) // 验证邮箱
	}
}
