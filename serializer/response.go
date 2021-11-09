package serializer

// Response 响应信息结构体
type Response struct {
	Code errCode     `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// NewResponse 响应信息构建
func NewResponse(code interface{}, data interface{}, msg string) Response {
	return Response{
		Code: code.(errCode),
		Data: data,
		Msg:  msg,
	}
}

// CodeResponse 错误码对应错误信息
func CodeResponse(errCode errCode) Response {
	return NewResponse(errCode, nil, Msg(errCode))
}

// SuccessResponse 通用成功信息
func SuccessResponse() Response {
	return CodeResponse(CodeSuccess)
}

// FailedResponse 通用失败信息，一般是客户端引起的失败
func FailedResponse() Response {
	return CodeResponse(CodeFailure)
}

// ErrorResponse 通用错误信息，一般是服务端错误
func ErrorResponse() Response {
	return CodeResponse(CodeError)
}

func ParamErrorResponse() Response {
	return CodeResponse(CodeParamErr)
}
