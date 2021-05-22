package user

import (
	"strconv"

	"acgfate/model/user"
)

// BaseInfoResponse 用户信息结构
type BaseInfoResponse struct {
	UID          uint64 `json:"uid"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	Mail         string `json:"mail"`
	MailVerified bool   `json:"mail_verify"`
	JoinTime     string `json:"join_time"`
	AccountState string `json:"account_state"`
	Sign         string `json:"sign"`
	Gender       string `json:"gender"`
	Credit       uint32 `json:"credit"`
	Birthday     string `json:"birthday"`
	Province     string `json:"province"`
	City         string `json:"city"`
}

// BuildBaseInfoResponse 构建用户信息响应
func BuildBaseInfoResponse(info *user.BaseInfo) BaseInfoResponse {
	joinTimeStr := strconv.FormatInt(info.JoinTime.Unix(), 10)
	accountStateStr := stateName(info.AccountState)
	genderStr := genderName(info.Gender)
	return BaseInfoResponse{
		UID:          info.UID,
		Username:     info.Username,
		Nickname:     info.Nickname,
		Mail:         info.Mail,
		MailVerified: info.MailVerified,
		JoinTime:     joinTimeStr,
		AccountState: accountStateStr,
		Sign:         info.Sign,
		Gender:       genderStr,
		Credit:       info.Credit,
		Birthday:     info.Birthday.Format("2006-01-02"),
		Province:     info.Province,
		City:         info.City,
	}
}

// genderName Gender number to name
func genderName(gender uint8) string {
	switch gender {
	case 1:
		return "男"
	case 2:
		return "女"
	case 3:
		return "其他"
	default:
		return "保密"
	}
}

// stateName Account State to name
func stateName(state uint8) string {
	switch state {
	case 0:
		return "正常"
	case 1:
		return "封禁"
	default:
		return "正常"
	}
}
