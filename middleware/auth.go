package middleware

import (
	"net/http"
	"strings"

	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authStr := c.Request.Header.Get("Authorization")
		msg := ""
		// 判断是否为空
		if authStr == "" {
			msg = "请求头Authorization为空"
		}
		// 格式化Authorization
		strParts := strings.SplitN(authStr, " ", 2)
		if len(strParts) != 2 || strParts[0] != "Bearer" {
			msg = "请求头Authorization格式错误"
		}
		// 解析判断是否正确
		res, err := utils.ParseToken(strParts[1])
		if err != nil {
			msg = "Token无效"
		}
		// 返回错误
		if msg != "" {
			c.JSON(http.StatusOK, gin.H{
				"code": sz.AccAuthErr,
				"msg":  msg,
			})
			log.Logger.Info(msg) // log
			c.Abort()
			return
		}
		CurrentUser(c, res.UID) // 保存当前用户信息到上下文
		c.Next()
	}
}

// CurrentUser 绑定当前用户
func CurrentUser(c *gin.Context, uid interface{}) {
	var acc model.Account
	if err := acc.BindAccount(uid); err != nil {
		log.Logger.Errorf(err.Error())
		return
	}
	c.Set("USER", &acc)
}
