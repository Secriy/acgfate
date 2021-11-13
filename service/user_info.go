package service

import (
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type UserInfoService struct{}

// Info 用户基本信息查询服务
func (service *UserInfoService) Info(c *gin.Context) (resp sz.Response) {
	user := model.CurrentUser(c)
	if user == nil {
		return sz.CodeResponse(sz.CodeAccAuthErr)
	}
	// 构建模型
	resp = sz.Success()
	resp.Data = sz.NewUser(user)
	return
}
