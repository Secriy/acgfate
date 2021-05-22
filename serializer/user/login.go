package user

import "acgfate/model/user"

// LoginResponse 登录信息返回
type LoginResponse struct {
	UID      uint64 `json:"uid"`      // UID
	Username string `json:"username"` // 用户名
	Token    string `json:"token"`    // Token
}

// BuildLoginResponse 登录信息返回构建
func BuildLoginResponse(info *user.BaseInfo, token string) LoginResponse {
	return LoginResponse{
		UID:      info.UID,
		Username: info.Username,
		Token:    token,
	}
}
