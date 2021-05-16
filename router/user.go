package router

import (
	v1 "acgfate/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	pubRouter := r.Group("user")
	{
		pubRouter.POST("register", v1.UserRegister)
		pubRouter.POST("login", v1.UserLogin)
	}

	// userAuthRouter := pubRouter.Use(middleware.JWTAuthRequired())
	// {
	// }
}
