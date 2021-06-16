package serializer

import (
	"strconv"

	"acgfate/model"
)

// BasicInfoResponse 用户信息结构
type BasicInfoResponse struct {
	UID      uint64 `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	JoinTime string `json:"join_time"`
	Level    int    `json:"level"`
	Exp      int    `json:"exp"`
	Sign     string `json:"sign"`
	Gender   string `json:"gender"`
	Credit   uint32 `json:"credit"`
	Birthday string `json:"birthday"`
	Province string `json:"province"`
	City     string `json:"city"`
}

// BuildBaseInfoResponse 构建用户信息响应
func BuildBaseInfoResponse(info *model.BasicInfo, username string) Response {
	joinTimeStr := strconv.FormatInt(info.JoinTime.Unix(), 10)
	genderStr := genderName(info.Gender)
	return BuildResponse(Success, BasicInfoResponse{
		UID:      info.UID,
		Username: username,
		Nickname: info.Nickname,
		JoinTime: joinTimeStr,
		Level:    model.GetLevelByExp(info.Exp),
		Exp:      info.Exp,
		Sign:     info.Sign,
		Gender:   genderStr,
		Credit:   info.Credit,
		Birthday: info.Birthday.Format("2006-01-02"),
		Province: info.Province,
		City:     info.City,
	}, Msg(Success))
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
