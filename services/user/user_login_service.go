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

func (service *LoginService) Login(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err != nil {
		return serializer.Error(5000, "账号或密码错误")
	}

	if !user.CheckPass(service.Password) {
		return serializer.Error(5000, "账号或密码错误")
	}

	token := RetJwt(user.UID)
	if token == "" {
		return serializer.Error(6000, "生成token失败")
	}

	return serializer.BuildResponse(200, serializer.BuildLoginResponse(&user, token), "登录成功")
}

func RetJwt(uid uint64) string {
	token, err := utils.GenToken(uid)
	if err != nil {
		return ""
	}
	return token
}
