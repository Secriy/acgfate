package services

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
func (service *UsernameService) CheckUsername() sz.Response {
	var dao database.UserDao
	if dao.IsExists(database.QUname, service.Username) {
		return sz.ErrResponse(sz.RegNameExist)
	}
	return sz.SuccessResponse()
}

// CheckEmail 判断邮箱是否被占用
func (service *EmailService) CheckEmail() sz.Response {
	var dao database.UserDao
	if dao.IsExists(database.QEmail, service.Email) {
		return sz.ErrResponse(sz.EmailExist)
	}
	return sz.SuccessResponse()
}
