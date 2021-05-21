package user

import (
	"fmt"

	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type MeService struct{}

// Me 用户信息查询服务
func (service *MeService) Me(c *gin.Context) sz.Response {
	uid := c.GetUint64("UID") // 当前用户UID
	// 绑定数据
	userInfo, err := model.GetUserInfo(uid)
	if err != nil {
		return sz.ErrResponse(sz.Error, "查询个人信息错误")
	}
	userPoints, err := model.GetUserPoint(uid)
	if err != nil {
		fmt.Println(err.Error())
		return sz.ErrResponse(sz.Error, "查询个人信息错误b")
	}
	// 构建模型
	return sz.BuildResponse(
		sz.Success,
		sz.BuildUserResponse(&userInfo, &userPoints),
		sz.GetResMsg(sz.Success),
	)
}
