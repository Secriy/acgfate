package user

import (
	"context"
	"fmt"
	"time"

	"acgfate/cache"
	"acgfate/log"
	"acgfate/model"
	"acgfate/model/user"
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
	var baseInfo user.BaseInfo
	if err := baseInfo.GetBaseInfo(c); err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.ErrResponse(sz.QueryDBErr)
	}
	// 生成验证码
	verifyCode := utils.GenerateCode(6)
	// 存储验证码
	mail := baseInfo.Mail
	key := fmt.Sprintf("MailVerifyCode:%s", mail)
	if err := cache.RDB.Set(ctx, key, verifyCode, 120*time.Second).Err(); err != nil { // 超时时间120s
		log.Logger.Errorf("%s: %s", sz.Msg(sz.CacheStoreErr), err)
		return sz.ErrResponse(sz.Error)
	}
	// 发送验证码
	err := utils.SendVerificationCode(mail, verifyCode)
	if err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.MailSendErr), err)
		return sz.ErrResponse(sz.Error)
	}

	log.Logger.Infof("发送验证码成功: %s", mail)

	return sz.SuccessResponse()
}

// Verify 校验验证码服务
func (service *MailVerifyService) Verify(c *gin.Context) sz.Response {
	ctx := context.Background()
	// 绑定用户模型
	var baseInfo user.BaseInfo
	if err := baseInfo.GetBaseInfo(c); err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.ErrResponse(sz.QueryDBErr)
	}
	// 获取缓存中的验证码
	mail := baseInfo.Mail
	key := fmt.Sprintf("MailVerifyCode:%s", mail)
	val, err := cache.RDB.Get(ctx, key).Result()
	// 判断是否发送验证码或过期
	if err == redis.Nil {
		return sz.ErrResponse(sz.VerifyCodeExpired)
	}
	// 判断验证码是否正确
	if val != service.Code {
		k := fmt.Sprintf("InputCount:%s", mail)
		if v, _ := cache.RDB.Get(ctx, k).Int(); v > 3 {
			cache.RDB.Del(ctx, key) // 输错三次
			cache.RDB.Del(ctx, k)
		}
		cache.RDB.Incr(ctx, k) // 输错一次加1
		return sz.ErrResponse(sz.VerifyCodeIncorrect)
	}
	// 写数据库
	uid := c.GetUint64("UID")
	sqlStr := "UPDATE user_base_info SET mail_verified = ? where uid = ? AND mail = ?"
	_, err = model.DB.Exec(sqlStr, true, uid, mail)
	if err != nil {
		log.Logger.Errorf("更新用户信息失败: %s", err)
		return sz.MsgResponse(sz.UpdateDBErr, "邮箱验证失败")
	}

	log.Logger.Infof("%s 验证成功", mail)

	return sz.SuccessResponse()
}
