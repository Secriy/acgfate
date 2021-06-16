package services

import (
	"acgfate/log"
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
	var acc model.Account
	// 检查账号是否存在 检查密码是否正确
	if err := acc.BindAccountByUsername(service.Username); err != nil || !acc.CheckPassword(service.Password) {
		return sz.MsgResponse(sz.Failure, "账号或密码错误")
	}
	// 生成用户Token
	token, err := utils.GenToken(acc.UID)
	if err != nil {
		log.Logger.Infof("生成token失败: %s", err)
		return sz.ErrResponse(sz.TokenGenerateErr)
	}

	log.Logger.Infof("登录成功: %d", acc.UID)

	return sz.BuildLoginResponse(&acc, token)
}
