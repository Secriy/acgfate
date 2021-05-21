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
	var userInfo model.UserInfo
	// 检查账号是否存在 && 检查密码是否正确
	sqlStr := "SELECT * FROM user_info where username = ?"
	err := model.DB.Get(&userInfo, sqlStr, service.Username)
	if err != nil {
		return sz.ErrResponse(sz.Failure, "账号或密码错误")
	}
	if !userInfo.CheckPassword(service.Password) {
		return sz.ErrResponse(sz.Failure, "账号或密码错误")
	}
	// 生成用户Token
	token, err := utils.GenToken(userInfo.UID)
	if err != nil {
		return sz.ErrResponse(sz.Error, "生成token失败")
	}

	return sz.BuildResponse(
		sz.Success,
		sz.BuildLoginResponse(&userInfo, token),
		"登录成功",
	)
}
