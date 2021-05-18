package serializer

import "acgfate/model"

// UserPointsResponse 用户点数信息
type UserPointsResponse struct {
	EXP   int64 `json:"exp"`
	Level uint8 `json:"level"`
	Coins uint  `json:"coins"`
}

// BuildUserPointsResponse 用户点数信息返回构建
func BuildUserPointsResponse(user *model.UserPoints) UserPointsResponse {
	return UserPointsResponse{
		EXP:   user.EXP,
		Level: model.FormatLevel(user.EXP),
		Coins: user.Coins,
	}
}
