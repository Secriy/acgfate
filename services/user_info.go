package services

import (
	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type InfoService struct{}

// Info 用户基本信息查询服务
func (service *InfoService) Info(c *gin.Context) sz.Response {
	// 绑定用户模型
	var basicInfo model.BasicInfo
	var username string
	username, err := basicInfo.CurrentBasicInfo(c)
	if err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.ErrResponse(sz.QueryDBErr)
	}

	// 构建模型
	return sz.BuildBaseInfoResponse(&basicInfo, username)
}
