package service

import (
	"acgfate/database"
	sz "acgfate/serializer"
)

type UsernameService struct {
	Username string `form:"username" binding:"required,alphanum,min=2,max=10"`
}

type EmailService struct {
	Email string `form:"email" binding:"required,email"`
}

// CheckUsername 判断用户名是否被占用
func (s *UsernameService) CheckUsername() sz.Response {
	var dao database.UserDao
	if _, err := dao.QueryByUname(s.Username); err == nil {
		return sz.CodeResponse(sz.CodeRegNameExist)
	}
	return sz.Success()
}

// CheckEmail 判断邮箱是否被占用
func (s *EmailService) CheckEmail() sz.Response {
	var dao database.UserDao
	if _, err := dao.QueryByEmail(s.Email); err == nil {
		return sz.CodeResponse(sz.CodeEmailExist)
	}
	return sz.Success()
}
