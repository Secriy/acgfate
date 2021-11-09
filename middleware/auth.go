package middleware

import (
	"net/http"

	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret, rdb string) gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", rdb, "", []byte(secret))
	// Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions("af-session", store)
}

// CurrentUser 绑定当前用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("uid")
		if uid != nil {
			var dao database.UserDao
			user, err := dao.QueryByUID(uid)
			if err == nil {
				c.Set("user", user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code": sz.AccAuthErr,
			"msg":  sz.Message(sz.AccAuthErr),
		})
		c.Abort()
	}
}
