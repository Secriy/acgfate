package user

import (
	"time"

	"acgfate/model"
	sz "acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
)

type MailVerifyCodeService struct{}

type MailVerifyService struct {
	Code string `json:"code" binding:"numeric,min=6,max=6"`
}

var (
	verifyCode string
	codeTime   int64
)

// SendCode 发送验证码
func (service MailVerifyCodeService) SendCode(c *gin.Context) sz.Response {
	// 生成验证码
	verifyCode = utils.GenerateCode(6)
	codeTime = time.Now().Unix()
	// 发送验证码
	err := utils.SendVerificationCode("secriy@qq.com", verifyCode)
	if err != nil {
		return sz.Err(sz.Error, "发送失败")
	}

	return sz.BuildResponse(
		sz.Success,
		nil,
		"发送成功",
		nil,
	)
}

// Verify 验证验证码
func (service MailVerifyService) Verify(c *gin.Context) sz.Response {
	var userInfo model.UserInfo
	// 绑定数据
	if err := model.DB.First(&userInfo, c.GetUint64("UID")).Error; err != nil {
		return sz.Err(sz.Error, "查询个人信息错误")
	}
	// 判断是否发送验证码
	if verifyCode == "" {
		return sz.Err(sz.Failure, "未发送验证码")
	}
	// 判断验证码是否过期
	if codeTime+120000 < time.Now().Unix() {
		return sz.Err(sz.Failure, "验证码已过期")
	}
	// 判断验证码是否正确
	if verifyCode != service.Code {
		return sz.Err(sz.Failure, "验证码不正确")
	}
	// 更新验证状态
	model.DB.Model(&userInfo).Update("mail_verify", true)

	return sz.BuildResponse(
		sz.Success,
		sz.BuildMailResponse(&userInfo),
		sz.GetResMsg(sz.Success),
		nil,
	)
}
