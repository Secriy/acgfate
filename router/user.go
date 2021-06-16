package router

import (
	"acgfate/api/http/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	// 公开路由组
	pub := r.Group("user")
	{
		pub.POST("register", v1.UserRegister) // 注册
		pub.POST("login", v1.UserLogin)       // 登录
	}

	// 需要鉴权的路由组
	auth := pub.Use(middleware.JWTAuthRequired())
	{
		auth.GET("info", v1.UserInfo)         // 获取个人信息
		auth.PUT("update", v1.UserInfoUpdate) // 更新个人基础信息
	}

	// 邮箱路由组
	mail := r.Group("email").Use(middleware.JWTAuthRequired())
	{
		mail.GET("verify", v1.MailSend)    // 发送验证码
		mail.POST("verify", v1.MailVerify) // 验证邮箱
	}
}
