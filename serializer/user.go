package serializer

import (
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
	UID        uint64             `json:"uid"`
	Username   string             `json:"username"`
	Nickname   string             `json:"nickname"`
	Mail       string             `json:"mail"`
	Avatar     string             `json:"avatar"`
	Gender     string             `json:"gender"`
	Birthday   string             `json:"birthday"`
	JoinTime   int64              `json:"join_time"`
	Silence    bool               `json:"silence"`
	UserPoints UserPointsResponse `json:"user_points"`
}

// BuildLoginResponse 登录信息返回构建
func BuildLoginResponse(user *model.UserInfo, token string) LoginResponse {
	return LoginResponse{
		UID:      user.UID,
		Username: user.Username,
		Token:    token,
	}
}

// BuildUserResponse 用户信息返回构建
func BuildUserResponse(user *model.User) UserResponse {
	return UserResponse{
		UID:        user.UserInfo.UID,
		Username:   user.Username,
		Nickname:   user.Nickname,
		Mail:       user.Mail,
		Avatar:     user.Avatar,
		Gender:     model.GenderFlags[int(user.Gender)],
		Birthday:   user.Birthday,
		JoinTime:   user.JoinTime.Unix(),
		Silence:    user.Silence,
		UserPoints: BuildUserPointsResponse(&user.UserPoints),
	}
}
