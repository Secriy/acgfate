package service

import (
	"acgfate/database"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserLoginService struct {
	Username string `json:"username" binding:"required,alphanum,min=2,max=10"`
	Password string `json:"password" binding:"required,ascii,min=8,max=16"`
}

// Login 用户登录服务
func (s *UserLoginService) Login(c *gin.Context) sz.Response {
	dao := new(database.UserDao)
	user, err := dao.QueryByUname(s.Username)
	if err != nil {
		zap.S().Debugf("登录失败: %e", err)
		return sz.CodeResponse(sz.CodeLoginIncorrect)
	}
	if !user.CheckPassword(s.Password) {
		return sz.CodeResponse(sz.CodeLoginIncorrect)
	}
	s.SetSession(c, user)

	zap.S().Infof("登录成功: %d", user.UID)

	return sz.Success()
}

// SetSession 保存 session
func (_ *UserLoginService) SetSession(c *gin.Context, user *model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("uid", user.UID)
	_ = s.Save()
}
