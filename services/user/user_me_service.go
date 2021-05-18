package user

import (
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type MeService struct{}

// Me 用户信息查询服务
func (service *MeService) Me(c *gin.Context) sz.Response {
	var userInfo model.UserInfo
	var userPoints model.UserPoints
	// 绑定数据
	if err := model.DB.First(&userInfo, c.GetUint64("UID")).Error; err != nil {
		return sz.ErrorResonse(sz.Error, "查询个人信息错误")
	}
	if err := model.DB.First(&userPoints, c.GetUint64("UID")).Error; err != nil {
		return sz.ErrorResonse(sz.Error, "查询个人信息错误")
	}
	// 构建模型
	user := model.User{
		UserInfo:   userInfo,
		UserPoints: userPoints,
	}

	return sz.BuildResponse(sz.Success, sz.BuildUserResponse(&user),
		sz.GetResMsg(sz.Success))
}
