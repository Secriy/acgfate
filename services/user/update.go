package user

import (
	"fmt"

	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

type UpdateService struct {
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=15"`
	Mail     string `json:"mail" binding:"omitempty,email"`
	Gender   uint8  `json:"gender" binding:"omitempty,min=1,max=4"`
	Birthday string `json:"birthday" binding:"omitempty"`
}

// Update 用户信息更新服务
func (service *UpdateService) Update(c *gin.Context) sz.Response {
	userInfo, err := model.GetUserInfo(c.GetUint64("UID"))
	if err != nil {
		return sz.ErrResponse(sz.Error, "获取当前用户失败")
	}
	// 判断是否修改邮箱
	var mailVerify bool
	switch {
	case userInfo.MailVerified == false:
		mailVerify = false
	case userInfo.Mail == service.Mail:
		mailVerify = true
	default:
		mailVerify = false
	}
	// 更新数据
	sql := "UPDATE user_info SET nickname = ?, mail = ?, mail_verify = ?, gender = ?, birthday = ? where uid = ?"
	_, err = model.DB.Exec(sql, service.Nickname, service.Mail, mailVerify, service.Gender, service.Birthday)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
	}
	return sz.BuildResponse(
		sz.Success,
		sz.BuildUserInfoResponse(&userInfo),
		sz.GetResMsg(sz.Success),
	)
}
