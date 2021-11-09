package service

import (
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type InfoService struct{}

// Info 用户基本信息查询服务
func (service *InfoService) Info(c *gin.Context) sz.Response {
	user := model.CurrentUser(c)
	if user == nil {
		return sz.ErrResponse(sz.AccAuthErr)
	}
	// 构建模型
	return sz.BuildUserResponse(user)
}
