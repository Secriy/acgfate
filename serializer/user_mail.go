package serializer

import "acgfate/model"

// MailResponse 登录信息返回
type MailResponse struct {
	Mail   string `json:"mail"`   // Email
	Verify string `json:"verify"` // 验证
}

// BuildMailResponse 构建邮箱返回信息
func BuildMailResponse(userInfo *model.UserInfo) MailResponse {
	isVerify := map[bool]string{
		true:  "已验证",
		false: "未验证",
	}
	return MailResponse{
		Mail:   userInfo.Mail,
		Verify: isVerify[userInfo.MailVerified],
	}
}
