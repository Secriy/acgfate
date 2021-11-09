package services

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
}

// Login 用户登录服务
func (service *LoginService) Login(c *gin.Context) sz.Response {
	var dao database.UserDao
	user, err := dao.QueryRow(database.QUname, service.Username)
	if err != nil {
		zap.S().Debugf("登录失败: %e", err)
		return sz.MsgResponse(sz.Failure, "账号或密码错误")
	}
	if !user.CheckPassword(service.Password) {
		return sz.MsgResponse(sz.Failure, "账号或密码错误")
	}
	service.SetSession(c, user)
	zap.S().Infof("登录成功: %d", user.UID)

	return sz.SuccessResponse()
}

// SetSession 保存 session
func (service *LoginService) SetSession(c *gin.Context, user *model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("uid", user.UID)
	_ = s.Save()
}
