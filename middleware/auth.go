package middleware

import (
	"net/http"
	"strings"

	sz "acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.AccAuthErr,
				"msg":  "Authorization为空",
			})
			c.Abort()
			return
		}
		token := strings.SplitN(authHeader, "Bearer ", 2)
		if token[0] != "" {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.AccAuthErr,
				"msg":  "Authorization格式不正确",
			})
			c.Abort()
			return
		}
		info, err := utils.ParseToken(token[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.AccAuthErr,
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
