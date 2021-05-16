package serializer

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// BuildResponse 响应信息构建
func BuildResponse(code int, data interface{}, msg string) Response {
	return Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

// Error 错误信息构建
func Error(code int, msg string) Response {
	return Response{
		Code: code,
		Data: nil,
		Msg:  msg,
	}
}
