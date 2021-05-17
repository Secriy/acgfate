package serializer

import (
	"time"

	"acgfate/model"
)

// LoginResponse 登录信息返回
type LoginResponse struct {
	UID      uint64 `json:"uid"`      // UID
	Username string `json:"username"` // 用户名
	Token    string `json:"token"`    // Token
}

// UserResponse 用户信息返回
type UserResponse struct {
	UID      uint64    `json:"uid"`
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	Mail     string    `json:"mail"`
	Avatar   string    `json:"avatar"`
	Gender   string    `json:"gender"`
	Level    uint8     `json:"level"`
	JoinTime time.Time `json:"join_time"`
	Silence  bool      `json:"silence"`
}

// BuildLoginResponse 登录信息返回构建
func BuildLoginResponse(user *model.User, token string) LoginResponse {
	return LoginResponse{
		UID:      user.UID,
		Username: user.Username,
		Token:    token,
	}
}

// BuildUserResponse 用户信息返回构建
func BuildUserResponse(user *model.User) UserResponse {
	return UserResponse{
		UID:      user.UID,
		Username: user.Username,
		Nickname: user.Nickname,
		Mail:     user.Mail,
		Avatar:   user.Avatar,
		Gender:   model.GenderFlags[int(user.Gender)],
		Level:    user.Level,
		JoinTime: user.JoinTime,
		Silence:  user.Silence,
	}
}
