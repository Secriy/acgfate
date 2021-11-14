package service

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserRegisterService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
	Nickname string `json:"nickname" binding:"required,min=2,max=15"`
	Email    string `json:"email" binding:"required,email,min=3,max=100"`
}

// Register 用户注册服务
func (s *UserRegisterService) Register() sz.Response {
	var user model.User
	dao := new(database.UserDao)
	// 判断用户名是否被占用
	if _, err := dao.QueryByUname(s.Username); err == nil {
		return sz.CodeResponse(sz.CodeRegNameExist)
	}
	// 判断邮箱是否被占用
	if _, err := dao.QueryByEmail(s.Email); err == nil {
		return sz.CodeResponse(sz.CodeEmailExist)
	}
	// 加密密码
	if err := user.SetPassword(s.Password); err != nil {
		zap.S().Errorf("%s: %s", sz.CodePasswdEncryptErr.String(), err)
		return sz.Error()
	}
	// 创建用户账号记录
	user.Username = s.Username
	user.Nickname = s.Nickname
	user.Email = s.Email
	if err := dao.Insert(&user); err != nil {
		zap.S().Errorf("创建用户失败: %s", err)
		return sz.Error()
	}
	return sz.Success()
}
