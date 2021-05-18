package user

import (
	"acgfate/model"
	sz "acgfate/serializer"
	"acgfate/utils"
)

type LoginService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
}

// Login 用户登录服务
func (service *LoginService) Login() sz.Response {
	var user model.UserInfo
	// 检查账号是否存在
	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err != nil {
		return sz.ErrorResonse(sz.AccAuthErr, "账号或密码错误")
	}
	// 检查密码是否正确
	if !user.CheckPass(service.Password) {
		return sz.ErrorResonse(sz.AccAuthErr, "账号或密码错误")
	}
	// 生成用户Token
	token, err := utils.GenToken(user.UID)
	if err != nil {
		return sz.ErrorResonse(sz.Error, "生成token失败")
	}

	return sz.BuildResponse(sz.Success, sz.BuildLoginResponse(&user, token), "登录成功")
}
