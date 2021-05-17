package user

import (
	"acgfate/model"
	"acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type UpdateService struct {
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=15"`
	Mail     string `json:"mail" binding:"omitempty,email"`
	Gender   uint8  `json:"gender" binding:"omitempty,min=1,max=4"`
	Birthday string `json:"birthday" binding:"omitempty"`
}

// Update 用户信息更新服务
func (service *UpdateService) Update(c *gin.Context) serializer.Response {
	user, err := model.GetUser(c.GetUint64("UID"))
	if err != nil {
		return serializer.Error(utils.Error, "获取当前用户失败")
	}
	model.DB.Model(&user.UserInfo).Updates(model.UserInfo{
		Nickname: service.Nickname,
		Mail:     service.Mail,
		Gender:   service.Gender,
		Birthday: service.Birthday,
	})

	return serializer.BuildResponse(utils.Success, serializer.BuildUserResponse(&user),
		utils.GetResMsg(utils.Success))
}
