package user

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"acgfate/cache"
	"acgfate/model"
	sz "acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type MailVerifyCodeService struct{}

type MailVerifyService struct {
	Code string `json:"code" binding:"numeric,min=6,max=6"`
}

// SendCode 发送验证码
func (service MailVerifyCodeService) SendCode(c *gin.Context) sz.Response {
	ctx := context.Background()
	// 生成验证码
	verifyCode := utils.GenerateCode(6)
	// 存储验证码
	key := "MailVerifyCode:" + strconv.FormatUint(c.GetUint64("UID"), 10)
	if err := cache.RDB.Set(ctx, key, verifyCode, 120*time.Second).Err(); err != nil {
		return sz.ErrResponse(sz.Error, "存储验证码失败")
	}
	// 发送验证码
	err := utils.SendVerificationCode("secriy@qq.com", verifyCode)
	if err != nil {
		return sz.ErrResponse(sz.Error, "发送失败")
	}

	return sz.BuildResponse(
		sz.Success,
		nil,
		"发送成功",
	)
}

// Verify 验证验证码
func (service MailVerifyService) Verify(c *gin.Context) sz.Response {
	ctx := context.Background()
	// 绑定数据
	userInfo, err := model.GetUserInfo(c.GetUint64("UID"))
	if err != nil {
		return sz.ErrResponse(sz.Error, "查询个人信息错误")
	}
	// 获取缓存中的验证码
	key := "MailVerifyCode:" + strconv.FormatUint(c.GetUint64("UID"), 10)
	val, err := cache.RDB.Get(ctx, key).Result()
	// 判断是否发送验证码
	if err == redis.Nil {
		return sz.ErrResponse(sz.Failure, "验证码已过期")
	}
	// 判断验证码是否正确
	if val != service.Code {
		return sz.ErrResponse(sz.Failure, "验证码不正确")
	}
	cache.RDB.Del(ctx, key)
	// 更新验证状态
	sql := "UPDATE user_info SET mail_verify = ? where uid = ?"
	_, err = model.DB.Exec(sql, true, userInfo.UID)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
	}

	return sz.BuildResponse(
		sz.Success,
		sz.BuildMailResponse(&userInfo),
		sz.GetResMsg(sz.Success),
	)
}
