package middleware

import (
	"net/http"

	"acgfate/model/user"
	"acgfate/serializer"
	"github.com/gin-gonic/gin"
)

// IsBanned 判断账号是否被封禁
func IsBanned() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo, err := user.GetUserInfoByID(c.GetUint64("UID"))
		var msg string
		switch {
		case err != nil:
			msg = err.Error()
		case userInfo.AccountState == user.StatusBanned:
			msg = "账号已被封禁"
		}
		if msg != "" {
			c.JSON(http.StatusOK, gin.H{
				"error": serializer.Failure,
				"msg":   msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
