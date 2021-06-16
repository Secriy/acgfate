package v1

import (
	"net/http"

	sz "acgfate/serializer"
	"acgfate/services"
	"github.com/gin-gonic/gin"
)

// MailSend 验证码邮件
// @Summary 验证码邮件发送
// @Description 验证码邮件发送接口
// @Tags Email
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Success 0 {object} serializer.Response "msg: "Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /email/verify [get]
func MailSend(c *gin.Context) {
	var form services.MailSendService
	res := form.Send(c)
	c.JSON(http.StatusOK, res)
}

// MailVerify 验证邮箱
// @Summary 验证邮件发送
// @Description 验证邮件发送接口
// @Tags Email
// @Accept application/json
// @Produce  application/json
// @Param Authorization header string true "用户令牌"
// @Param form body services.MailVerifyService true "验证码"
// @Success 0 {object} serializer.Response "msg: "Success"
// @Failure 30001 {object} serializer.Response "msg: 参数错误"
// @Router /email/verify [post]
func MailVerify(c *gin.Context) {
	var form services.MailVerifyService
	if err := c.ShouldBind(&form); err == nil {
		res := form.Verify(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, sz.ParmErr())
	}
}
