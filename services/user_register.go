package services

import (
	"time"

	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
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
	var acc model.Account
	// 判断用户名是否被占用
	if Exist("accounts", "username", service.Username) {
		return sz.ErrResponse(sz.RegNameExist)
	}
	// 判断邮箱是否被占用
	if Exist("accounts", "email", service.Email) {
		return sz.ErrResponse(sz.EmailExist)
	}
	// 加密密码
	if err := acc.SetPassword(service.Password); err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.PasswdEncryptErr), err)
		return sz.ErrResponse(sz.PasswdEncryptErr)
	}
	// 创建用户账号记录
	accSQL := "INSERT INTO accounts (username, password, email) VALUES (?,?,?)"
	_, err := model.DB.Exec(accSQL, service.Username, acc.Password, service.Email)
	if err != nil {
		log.Logger.Errorf("创建用户失败: %s", err)
		return sz.MsgResponse(sz.InsertDBErr, "创建用户失败")
	}
	// 创建用户基础信息记录
	infoSQL := "INSERT INTO  user_basic_info (nickname,join_time) VALUES (?,?)"
	_, err = model.DB.Exec(infoSQL, service.Nickname, time.Now())
	if err != nil {
		log.Logger.Errorf("创建用户失败: %s", err)
		return sz.MsgResponse(sz.InsertDBErr, "创建用户失败")
	}

	return sz.SuccessResponse()
}
