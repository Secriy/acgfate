package user

import (
	"time"

	"acgfate/model"
	sz "acgfate/serializer"
)

type RegisterService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
	Nickname string `json:"nickname" binding:"required,min=2,max=15"`
	Mail     string `json:"mail" binding:"required,email"`
}

// Register 用户注册服务
func (service RegisterService) Register() sz.Response {
	var userPoints model.UserPoints
	var userInfo = model.UserInfo{
		Username: service.Username,
		Password: service.Password,
		Nickname: service.Nickname,
		Mail:     service.Mail,
		JoinTime: time.Now(),
	}
	// 判断用户名是否已经存在
	if err := model.DB.Where("username = ?", service.Username).First(&userInfo).Error; err == nil {
		return sz.ErrorResponse(sz.AccCreateErr, "用户名已被他人使用")
	}
	// 加密密码
	if err := userInfo.SetPassword(service.Password); err != nil {
		return sz.ErrorResponse(
			sz.CodeEncryptError,
			"密码加密失败",
		)
	}
	// 创建用户
	if err := model.DB.Create(&userInfo).Error; err != nil {
		return sz.ErrorResponse(sz.DatabaseErr, "创建用户失败")
	}
	if err := model.DB.Create(&userPoints).Error; err != nil {
		return sz.ErrorResponse(sz.DatabaseErr, "创建用户失败")
	}
	// 构建模型
	user := model.User{
		UserInfo:   userInfo,
		UserPoints: userPoints,
	}

	return sz.BuildResponse(200, sz.BuildUserResponse(&user), sz.GetResMsg(sz.Success))
}
