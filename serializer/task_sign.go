package serializer

type TaskSignResponse struct {
	Reward string `json:"reward"`
}

// BuildTaskSignResponse 签到信息返回
func BuildTaskSignResponse(text string) Response {
	return BuildResponse(Success,
		TaskSignResponse{
			Reward: text,
		},
		Msg(Success))
}
