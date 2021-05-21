package router

import (
	v12 "acgfate/api/http/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

func InitWordsRouter(r *gin.RouterGroup) {
	pubGroup := r.Group("words")
	{
		pubGroup.GET(":wid", v12.WordsGet)
	}
	authGroup := pubGroup.Use(
		middleware.JWTAuthRequired(),
		middleware.IsVerified(),
		middleware.IsSilence(),
		middleware.IsBanned(),
	)
	{
		authGroup.POST("post", v12.WordsPost)
	}
}
