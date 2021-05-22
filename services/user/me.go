package user

import (
	"acgfate/log"
	"acgfate/model/user"
	sz "acgfate/serializer"
	suser "acgfate/serializer/user"
	"github.com/gin-gonic/gin"
)

type MeService struct{}

// Info 用户基本信息查询服务
func (service *MeService) Info(c *gin.Context) sz.Response {
	// 绑定用户模型
	var baseInfo user.BaseInfo
	if err := baseInfo.GetBaseInfo(c); err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.ErrResponse(sz.QueryDBErr)
	}

	// 构建模型
	return sz.BuildResponse(
		sz.Success,
		suser.BuildBaseInfoResponse(&baseInfo),
		sz.Msg(sz.Success),
	)
}
