package serializer

import (
	"acgfate/model"
)

// LoginResponse 登录信息返回
type LoginResponse struct {
	UID   uint64 `json:"uid"`   // UID
	Token string `json:"token"` // Token
}

// BuildLoginResponse 登录信息返回构建
func BuildLoginResponse(acc *model.Account, token string) Response {
	return BuildResponse(Success, LoginResponse{
		UID:   acc.UID,
		Token: token,
	}, Msg(Success))
}
