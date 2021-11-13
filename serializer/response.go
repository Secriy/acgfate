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

// DataResponse 获取错误码及输入数据的返回
func DataResponse(e errCode, data interface{}) Response {
	return NewResponse(e, data, e.String())
}

// CodeResponse 获取错误码对应的返回
func CodeResponse(e errCode) Response {
	return DataResponse(e, nil)
}

// Success 通用成功信息
func Success() Response {
	return CodeResponse(CodeSuccess)
}

// Failure 通用失败信息，一般是客户端引起的失败
func Failure() Response {
	return CodeResponse(CodeFailure)
}

// Error 通用错误信息，一般是服务端错误
func Error() Response {
	return CodeResponse(CodeError)
}

// ParamError 用户输入参数错误
func ParamError() Response {
	return CodeResponse(CodeParamErr)
}
