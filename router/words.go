package router

import (
	v1 "acgfate/api/http/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

// InitWordsRouter 初始化文字相关路由
func InitWordsRouter(r *gin.RouterGroup) {
	pub := r.Group("words")
	{
		pub.GET(":wid", v1.WordsGet)
	}
	auth := pub.Use(
		middleware.JWTAuthRequired(),
		middleware.IsVerified(),
		middleware.IsBanned(),
	)
	{
		auth.POST("post", v1.WordsPost)
	}
}
