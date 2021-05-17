package user

import (
	"time"

	"acgfate/model"
	"acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type RegisterService struct {
	Username string `json:"username" form:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" form:"password" binding:"required,ascii,min=8,max=16"`
	Nickname string `json:"nickname" form:"nickname" binding:"required,min=1,max=15"`
	Mail     string `json:"mail" form:"mail" binding:"required,email"`
}

// Register 用户注册服务
func (service RegisterService) Register(c *gin.Context) serializer.Response {
	user := model.User{
		Username: service.Username,
		Password: service.Password,
		Nickname: service.Nickname,
		Mail:     service.Mail,
		Level:    0,
		JoinTime: time.Now(),
		Silence:  false,
	}
	// 判断用户名是否已经存在
	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err == nil {
		return serializer.Error(utils.AccCreateErr, "用户名已被他人使用")
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Error(utils.DatabaseErr, "注册失败")
	}

	return serializer.BuildResponse(200, serializer.BuildUserResponse(&user), utils.GetResMsg(utils.Success))
}
