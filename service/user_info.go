package service

import (
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type UserInfoService struct{}

// UserInfo 用户基本信息查询服务
func (service *UserInfoService) UserInfo(c *gin.Context) sz.Response {
	user := model.CurrentUser(c)
	if user == nil {
		return sz.CodeResponse(sz.CodeAccAuthErr)
	}
	// 构建模型
	return sz.BuildUserResponse(user)
}
