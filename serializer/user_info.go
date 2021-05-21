package serializer

import (
	"strconv"

	"acgfate/model"
)

// UserResponse 用户返回
type UserResponse struct {
	UserInfoResponse
	UserPointsResponse
}

// LoginResponse 登录信息返回
type LoginResponse struct {
	UID      uint64 `json:"uid"`      // UID
	Username string `json:"username"` // 用户名
	Token    string `json:"token"`    // Token
}

// UserInfoResponse 用户信息返回
type UserInfoResponse struct {
	UID      uint64 `json:"uid"`
	JoinTime string `json:"join_time"`
	Username string `json:"username"`
	Nickname string `db:"nickname"`
	Mail     string `db:"mail"`
	Status   string `db:"status"`
	Avatar   string `db:"avatar"`
	Sign     string `db:"sign"`
	Gender   string `db:"gender"`
	Birthday string `db:"birthday"`
}

// BuildUserResponse 用户全部信息返回
func BuildUserResponse(userInfo *model.UserInfo, userPoints *model.UserPoints) UserResponse {
	return UserResponse{
		UserInfoResponse:   BuildUserInfoResponse(userInfo),
		UserPointsResponse: BuildUserPointsResponse(userPoints),
	}
}

// BuildLoginResponse 登录信息返回构建
func BuildLoginResponse(user *model.UserInfo, token string) LoginResponse {
	return LoginResponse{
		UID:      user.UID,
		Username: user.Username,
		Token:    token,
	}
}

// BuildUserInfoResponse 用户信息返回构建
func BuildUserInfoResponse(user *model.UserInfo) UserInfoResponse {
	return UserInfoResponse{
		UID:      user.UID,
		Username: user.Username,
		Nickname: user.Nickname,
		Mail:     user.Mail,
		Status:   model.GetStatus(user.Status),
		Avatar:   user.Avatar,
		Gender:   model.GetGender(user.Gender),
		Sign:     user.Sign.String,
		Birthday: user.Birthday.String,
		JoinTime: strconv.FormatInt(user.CreatedAt.Unix(), 10),
	}
}
