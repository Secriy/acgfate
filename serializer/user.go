package serializer

import (
	"time"

	"acgfate/model"
)

// User 用户信息结构
type User struct {
	UID      int64  `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	State    uint8  `json:"state"`
}

// UserInfo 用户个人信息
type UserInfo struct {
	UID      int64     `json:"uid"`
	Gender   uint8     `json:"gender"`
	Sign     string    `json:"sign"`
	Birthday time.Time `json:"birthday"`
	Province string    `json:"province"`
	City     string    `json:"city"`
}

// NewUser 构建用户信息
func NewUser(user *model.User) User {
	return User{
		UID:      user.UID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		State:    user.State,
	}
}

// NewUserInfo 构建用户信息
func NewUserInfo(userInfo *model.UserInfo) UserInfo {
	return UserInfo{
		UID:      userInfo.UID,
		Gender:   userInfo.Gender,
		Sign:     userInfo.Sign,
		Birthday: userInfo.Birthday,
		Province: userInfo.Province,
		City:     userInfo.City,
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
