package serializer

// Response 响应信息结构体
type Response struct {
	Code interface{} `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// BuildResponse 响应信息构建
func BuildResponse(code interface{}, data interface{}, msg string) Response {
	return Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

// MsgResponse 自定义消息错误返回
func MsgResponse(code interface{}, msg string) Response {
	return BuildResponse(code, nil, msg)
}
