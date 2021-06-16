package services

import (
	"context"
	"fmt"
	"time"

	"acgfate/cache"
	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
	"acgfate/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type MailSendService struct{}

type MailVerifyService struct {
	Code string `json:"code" binding:"required,numeric,min=6,max=6"`
}

// Send 发送验证码
func (service *MailSendService) Send(c *gin.Context) sz.Response {
	ctx := context.Background()
	// 绑定用户模型
	acc := model.CurrentUser(c)
	// 判断是否已经验证
	if acc.Verified {
		return sz.ErrResponse(sz.VerifyAlready)
	}
	// 生成验证码
	verifyCode := utils.GenerateCode(6)
	// 存储验证码
	email := acc.Email
	key := fmt.Sprintf("MailVerifyCode:%s", email)
	if err := cache.RDB.Set(ctx, key, verifyCode, 120*time.Second).Err(); err != nil { // 超时时间120s
		log.Logger.Errorf("%s: %s", sz.Msg(sz.CacheStoreErr), err)
		return sz.ErrResponse(sz.Error)
	}
	// 发送验证码
	err := utils.SendVerificationCode(email, verifyCode)
	if err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.MailSendErr), err)
		return sz.ErrResponse(sz.Error)
	}

	log.Logger.Infof("发送验证码成功: %s", email)

	return sz.SuccessResponse()
}

// Verify 校验验证码服务
func (service *MailVerifyService) Verify(c *gin.Context) sz.Response {
	ctx := context.Background()
	// 绑定用户模型
	acc := model.CurrentUser(c)
	// 判断是否已经验证
	if acc.Verified {
		return sz.ErrResponse(sz.VerifyAlready)
	}
	// 获取缓存中的验证码
	email := acc.Email
	key := fmt.Sprintf("MailVerifyCode:%s", email)
	val, err := cache.RDB.Get(ctx, key).Result()
	// 判断是否发送验证码或过期
	if err == redis.Nil {
		return sz.ErrResponse(sz.VerifyCodeExpired)
	}
	// 判断验证码是否正确
	k := fmt.Sprintf("InputCount:%s", email)
	if val != service.Code {
		if v, _ := cache.RDB.Get(ctx, k).Int(); v > 3 {
			cache.RDB.Del(ctx, key) // 输错三次
			cache.RDB.Del(ctx, k)
		}
		cache.RDB.Incr(ctx, k) // 输错一次加1
		return sz.ErrResponse(sz.VerifyCodeIncorrect)
	}
	cache.RDB.Del(ctx, k)
	// 写数据库
	sqlStr := "UPDATE accounts SET verified = ? where uid = ? AND email = ?"
	_, err = model.DB.Exec(sqlStr, true, acc.UID, email)
	if err != nil {
		log.Logger.Errorf("更新用户信息失败: %s", err)
		return sz.MsgResponse(sz.UpdateDBErr, "邮箱验证失败")
	}

	log.Logger.Infof("%s 验证成功", email)

	return sz.SuccessResponse()
}
