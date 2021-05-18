package serializer

const (
	Success = 0
	Error   = 50000

	ParamErr = 30001

	AccErr       = 40000
	AccAuthErr   = 40001
	AccCreateErr = 40002
	AccSilence   = 40009

	WordsPostErr = 41001

	DatabaseErr = 50002
)

var ResMsgFlags = map[int]string{
	Success:      "Success",
	Error:        "Fail",
	WordsPostErr: "发布失败",
}

func GetResMsg(code int) string {
	msg, ok := ResMsgFlags[code]
	if ok {
		return msg
	}
	return ResMsgFlags[Error]
}

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

// ErrorResonse 错误信息构建
func ErrorResonse(code int, msg string) Response {
	return Response{
		Code: code,
		Data: nil,
		Msg:  msg,
	}
}

// ErrorMsg 错误返回
func ErrorMsg(code int, err error) Response {
	return Response{
		Code: code,
		Data: nil,
		Msg:  err.Error(),
	}
}
