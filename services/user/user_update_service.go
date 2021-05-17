package user

import (
	"acgfate/model"
	"acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type UpdateService struct {
	Nickname string `json:"nickname" form:"nickname" binding:"max=15"`
	Mail     string `json:"mail" form:"mail" binding:"email"`
	Gender   uint8  `json:"gender"`
	Birthday string `json:"birthday"`
}

// Update 用户信息更新服务
func (service *UpdateService) Update(c *gin.Context) serializer.Response {
	user, err := model.GetUser(c.GetUint64("UID"))
	if err != nil {
		return serializer.Error(utils.Error, "获取当前用户失败")
	}
	model.DB.Model(&user).Updates(model.User{
		Nickname: service.Nickname,
		Mail:     service.Mail,
		Gender:   service.Gender,
	})

	return serializer.BuildResponse(utils.Success, serializer.BuildUserResponse(&user), utils.GetResMsg(utils.Success))
}
