package points

import (
	"fmt"
	"time"

	"acgfate/model"
	sz "acgfate/serializer"
	"github.com/gin-gonic/gin"
)

const (
	expReword   = 10 // 经验值奖励
	coinsReward = 2  // 硬币奖励
)

type SignService struct{}

// DoSign 签到服务
func (service *SignService) DoSign(c *gin.Context) sz.Response {
	user, err := model.GetUser(c.GetUint64("UID"))
	if err != nil {
		return sz.ErrorResponse(sz.Error, "获取当前用户失败")
	}
	if fmt.Sprint(user.SignTime.Date()) == fmt.Sprint(time.Now().Date()) {
		return sz.ErrorResponse(sz.SignErr, "今天已经签到过了")
	}
	// 更新用户点数
	model.DB.Model(&user.UserPoints).Updates(model.UserPoints{
		Coins:    user.Coins + coinsReward,
		EXP:      user.EXP + expReword,
		SignTime: time.Now(),
	})

	return sz.BuildResponse(sz.Success, sz.BuildUserPointsResponse(&user.UserPoints),
		sz.GetResMsg(sz.Success))
}
