package service

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
	Nickname string `json:"nickname" binding:"required,min=2,max=15"`
	Email    string `json:"email" binding:"required,email"`
}

// Register 用户注册服务
func (service *RegisterService) Register() sz.Response {
	var user model.User

	var dao database.UserDao
	// 判断用户名是否被占用
	if _, err := dao.QueryByUname(service.Username); err == nil {
		return sz.ErrResponse(sz.RegNameExist)
	}
	// 判断邮箱是否被占用
	if _, err := dao.QueryByEmail(service.Email); err == nil {
		return sz.ErrResponse(sz.EmailExist)
	}
	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		zap.S().Errorf("%s: %s", sz.Message(sz.PasswdEncryptErr), err)
		return sz.ErrResponse(sz.PasswdEncryptErr)
	}
	// 创建用户账号记录
	user.Username = service.Username
	user.Nickname = service.Nickname
	user.Email = service.Email
	err := dao.InsertRow(user)
	if err != nil {
		zap.S().Errorf("创建用户失败: %s", err)
		return sz.MsgResponse(sz.InsertDBErr, "创建用户失败")
	}
	return sz.SuccessResponse()
}
