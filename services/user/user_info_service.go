package user

import (
	"acgfate/model"
	"acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type MeService struct{}

// Me 用户信息查询服务
func (service *MeService) Me(c *gin.Context) serializer.Response {
	var user model.User
	// 绑定数据
	err := model.DB.First(&user, c.GetUint64("UID")).Error
	if err != nil {
		return serializer.Error(utils.Error, "查询个人信息错误")
	}

	return serializer.BuildResponse(utils.Success, serializer.BuildUserResponse(&user), utils.GetResMsg(utils.Success))
}
