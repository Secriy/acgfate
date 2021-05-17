package middleware

import (
	"net/http"

	"acgfate/model"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

func CheckSilence() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := model.GetUser(c.GetUint64("UID"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": utils.Error,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		if user.Silence {
			c.JSON(http.StatusOK, gin.H{
				"code": utils.AccSilence,
				"msg":  "账号已被禁言",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
