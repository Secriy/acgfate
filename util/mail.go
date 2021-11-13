package util

import (
	"crypto/tls"
	"fmt"

	"acgfate/config"
	"gopkg.in/gomail.v2"
)

// SendVerificationCode send email which include verification code
func SendVerificationCode(conf *config.EmailConfig, email string, code string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Sender) // 发件人
	m.SetHeader("To", email)         // 收件人
	m.SetHeader("Subject", "邮箱验证")
	m.SetBody("text/html", fmt.Sprintf("<h1>验证码：%s</h1>", code)) // mail content

	d := gomail.NewDialer(conf.Smtp, conf.Port, conf.Sender, conf.Passwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = d.DialAndSend(m) // send
	return
}
