package services

import (
	sz "acgfate/serializer"
)

type UsernameService struct {
	Username string `form:"username" binding:"required,alphanum,min=2,max=10"`
}

type EmailService struct {
	Email string `form:"email" binding:"required,email"`
}

func (service *UsernameService) CheckUsername() sz.Response {
	// 判断用户名是否被占用
	if Exist("accounts", "username", service.Username) {
		return sz.ErrResponse(sz.RegNameExist)
	}

	return sz.SuccessResponse()
}

func (service *EmailService) CheckEmail() sz.Response {
	// 判断邮箱是否被占用
	if Exist("accounts", "email", service.Email) {
		return sz.ErrResponse(sz.EmailExist)
	}
	return sz.SuccessResponse()
}
