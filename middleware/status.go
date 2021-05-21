package middleware

import (
	"net/http"

	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

func IsSilence() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := model.GetUserInfo(c.GetUint64("UID"))
		var msg string
		switch {
		case err != nil:
			msg = err.Error()
		case user.Status == model.StatusSilenced:
			msg = "账号已被禁言"
		}
		if msg != "" {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.Error,
				"msg":  msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func IsBanned() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := model.GetUserInfo(c.GetUint64("UID"))
		var msg string
		switch {
		case err != nil:
			msg = err.Error()
		case user.Status == model.StatusBanned:
			msg = "账号已被封禁"
		}
		if msg != "" {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.Error,
				"msg":  msg,
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
		user, err := model.GetUserInfo(c.GetUint64("UID"))
		var msg string
		switch {
		case err != nil:
			msg = err.Error()
		case user.MailVerified == false:
			msg = "账号未验证"
		}
		if msg != "" {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.Error,
				"msg":  msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
