package user

import (
	"acgfate/model"
	"acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	Username string `json:"username" form:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" form:"password" binding:"required,ascii,min=8,max=16"`
}

// Login 用户登录服务
func (service *LoginService) Login(c *gin.Context) serializer.Response {
	var user model.User
	// 检查账号是否存在
	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err != nil {
		return serializer.Error(utils.AccAuthErr, "账号或密码错误")
	}
	// 检查密码是否正确
	if !user.CheckPass(service.Password) {
		return serializer.Error(utils.AccAuthErr, "账号或密码错误")
	}
	// 生成用户Token
	token, err := utils.GenToken(user.UID)
	if err != nil {
		return serializer.Error(utils.Error, "生成token失败")
	}

	return serializer.BuildResponse(utils.Success, serializer.BuildLoginResponse(&user, token), "登录成功")
}
