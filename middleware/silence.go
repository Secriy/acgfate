package middleware

import (
	"net/http"

	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

func CheckSilence() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := model.GetUser(c.GetUint64("UID"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.Error,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		if user.Silence {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.AccSilence,
				"msg":  "账号已被禁言",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}