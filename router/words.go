package router

import (
	v1 "acgfate/api/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

func InitWordsRouter(r *gin.RouterGroup) {
	pubGroup := r.Group("words")
	{
		pubGroup.GET(":wid", v1.WordsGet)
	}
	authGroup := pubGroup.Use(middleware.JWTAuthRequired(), middleware.IsSilence(), middleware.IsMailVerify())
	{
		authGroup.POST("post", v1.WordsPost)
	}
}
