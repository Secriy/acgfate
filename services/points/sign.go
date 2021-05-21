package points

import (
	"fmt"
	"time"

	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

const (
	expReword   = 5 // 经验值奖励
	coinsReward = 2 // 硬币奖励
)

type SignService struct{}

// DoSign 签到服务
func (service *SignService) DoSign(c *gin.Context) sz.Response {
	userPoint, err := model.GetUserPoint(c.GetUint64("UID"))
	if err != nil {
		fmt.Println(err.Error())
		return sz.ErrResponse(sz.Error, "获取当前用户失败")
	}
	// 时间戳转时间
	if fmt.Sprint(userPoint.SignTime.Time.Date()) == fmt.Sprint(time.Now().Date()) {
		return sz.ErrResponse(sz.Failure, "今天已经签到过了")
	}
	// 更新用户点数
	sql := "UPDATE user_point SET coins = ?, exp = ?, sign_time = ? where uid = ?"
	userPoint.Coins += coinsReward // +硬币
	userPoint.EXP += expReword     // +经验
	_, err = model.DB.Exec(sql, userPoint.Coins, userPoint.EXP, time.Now(), userPoint.UID)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
	}

	return sz.BuildResponse(
		sz.Success,
		sz.BuildUserPointsResponse(&userPoint),
		sz.GetResMsg(sz.Success),
	)
}
