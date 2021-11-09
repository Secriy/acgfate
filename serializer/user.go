package serializer

import (
	"acgfate/model"
)

// UserResponse 用户信息结构
type UserResponse struct {
	UID      uint64 `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	State    uint8  `json:"state"`
}

// BuildUserResponse 构建用户信息响应
func BuildUserResponse(user *model.User) Response {
	return BuildResponse(Success, UserResponse{
		UID:      user.UID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		State:    user.State,
	}, Message(Success))
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
