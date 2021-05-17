package middleware

import (
	"net/http"

	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": utils.AccAuthErr,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		info, err := utils.ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": utils.Error,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		// 保存UID到上下文
		c.Set("UID", info.UID)
		c.Next()
	}
}
