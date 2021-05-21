package middleware

import (
	"net/http"
	"strings"

	"acgfate/config"
	sz "acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authStr := c.Request.Header.Get("Authorization")
		// 判断是否为空
		if authStr == "" {
			msg := "请求头Authorization为空"
			c.JSON(http.StatusOK, gin.H{
				"code": sz.AccAuthErr,
				"msg":  msg,
			})
			config.SugarLogger.Info(msg) // log
			c.Abort()
			return
		}
		// 格式化Authorization
		strParts := strings.SplitN(authStr, " ", 2)
		if len(strParts) != 2 || strParts[0] != "Bearer" {
			msg := "请求头Authorization格式错误"
			c.JSON(http.StatusOK, gin.H{
				"code": sz.AccAuthErr,
				"msg":  msg,
			})
			config.SugarLogger.Info(msg) // log
			c.Abort()
			return
		}
		// 解析判断是否正确
		res, err := utils.ParseToken(strParts[1])
		if err != nil {
			msg := "Token无效"
			c.JSON(http.StatusOK, gin.H{
				"code": sz.AccAuthErr,
				"msg":  msg,
			})
			config.SugarLogger.Info(msg) // log
			c.Abort()
			return
		}
		// 保存UID到上下文
		c.Set("UID", res.UID)
		c.Next()
	}
}
