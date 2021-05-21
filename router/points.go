package router

import (
	v12 "acgfate/api/http/v1"
	"acgfate/middleware"
	"github.com/gin-gonic/gin"
)

func InitPointsRouter(r *gin.RouterGroup) {
	// 需要鉴权的路由组
	group := r.Group("").Use(middleware.JWTAuthRequired())
	group.GET("sign", v12.PointsSign)
	// {
	// 	authGroup.GET("sign", v1.UserMe)          // 获取个人信息
	// 	authGroup.PUT("update", v1.UpdateService) // 更新个人信息
	// }
}
