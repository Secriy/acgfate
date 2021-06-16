package services

import (
	"context"
	"fmt"
	"time"

	"acgfate/cache"
	"acgfate/log"
	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

const (
	expReward   = 5 // 经验值奖励
	coinsReward = 2 // 硬币奖励
)

type SignService struct{}

// DoSign 签到服务
func (service *SignService) DoSign(c *gin.Context) sz.Response {
	ctx := context.Background()
	// 获取当前用户信息
	var baseInfo model.BasicInfo
	_, err := baseInfo.CurrentBasicInfo(c)
	if err != nil {
		log.Logger.Errorf("%s: %s", sz.Msg(sz.QueryDBErr), err)
		return sz.MsgResponse(sz.Error, "获取当前用户失败")
	}
	// 判断当天是否已经签到
	_, err = cache.RDB.Get(ctx, fmt.Sprintf("sign:%d:%s", baseInfo.UID, time.Now().Format("2006-01-02"))).Result()
	if err == nil {
		return sz.MsgResponse(sz.Failure, "今天已经签到过了")
	}
	// 存储签到信息到Redis
	cache.RDB.Set(ctx, fmt.Sprintf("sign:%d:%s", baseInfo.UID, time.Now().Format("2006-01-02")), true, time.Hour*100)
	// 更新用户点数
	sql := "UPDATE user_basic_info SET exp=exp+? where uid = ?"
	_, err = model.DB.Exec(sql, expReward, baseInfo.UID)
	if err != nil {
		log.Logger.Errorf("更新用户点数失败: %s", err)
		return sz.MsgResponse(sz.Error, "更新用户点数失败")
	}
	baseInfo.Exp += expReward

	msg := fmt.Sprintf("经验+%d", expReward)

	return sz.BuildTaskSignResponse(msg)
}
