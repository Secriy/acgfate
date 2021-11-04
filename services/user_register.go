package services

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"acgfate/utils/logger"
	_ "github.com/gin-gonic/gin"
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
	if dao.IsExists(database.QUname, service.Username) {
		return sz.ErrResponse(sz.RegNameExist)
	}
	// 判断邮箱是否被占用
	if dao.IsExists(database.QEmail, service.Email) {
		return sz.ErrResponse(sz.EmailExist)
	}
	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logger.Logger.Errorf("%s: %s", sz.Msg(sz.PasswdEncryptErr), err)
		return sz.ErrResponse(sz.PasswdEncryptErr)
	}
	// 创建用户账号记录
	user.Username = service.Username
	user.Nickname = service.Nickname
	user.Email = service.Email
	err := dao.InsertRow(user)
	if err != nil {
		logger.Logger.Errorf("创建用户失败: %s", err)
		return sz.MsgResponse(sz.InsertDBErr, "创建用户失败")
	}
	return sz.SuccessResponse()
}
