package user

import (
	"acgfate/log"
	"acgfate/model"
	"acgfate/model/user"
	sz "acgfate/serializer"
	suser "acgfate/serializer/user"
	"acgfate/utils"
)

type LoginService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
}

// Login 用户登录服务
func (service *LoginService) Login() sz.Response {
	var baseInfo user.BaseInfo
	// 检查账号是否存在
	sqlStr := "SELECT * FROM user_base_info where username = ?"
	err := model.DB.Get(&baseInfo, sqlStr, service.Username)
	if err != nil {
		log.Logger.Infof("账号或密码错误: %s", err)
		return sz.MsgResponse(sz.Failure, "账号或密码错误")
	}
	// 检查密码是否正确
	if !baseInfo.CheckPassword(service.Password) {
		log.Logger.Infof("账号或密码错误: %s", err)
		return sz.MsgResponse(sz.Failure, "账号或密码错误")
	}
	// 生成用户Token
	token, err := utils.GenToken(baseInfo.UID)
	if err != nil {
		log.Logger.Infof("生成token失败: %s", err)
		return sz.ErrResponse(sz.TokenGenerateErr)
	}

	log.Logger.Infof("登录成功: %d", baseInfo.UID)

	return sz.BuildResponse(
		sz.Success,
		suser.BuildLoginResponse(&baseInfo, token),
		"登录成功",
	)
}
