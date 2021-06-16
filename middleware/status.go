package middleware

import (
	"net/http"

	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

// IsBanned 判断账号是否被封禁
func IsBanned() gin.HandlerFunc {
	return func(c *gin.Context) {
		acc := model.CurrentUser(c)
		if acc.State == model.StatusBanned {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.Failure,
				"msg":  "账号已被封禁",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// IsVerified 判断账号是否验证
func IsVerified() gin.HandlerFunc {
	return func(c *gin.Context) {
		acc := model.CurrentUser(c)
		if !acc.Verified {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.Failure,
				"msg":  "账号邮箱未验证",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
