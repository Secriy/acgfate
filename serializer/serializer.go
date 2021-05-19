package serializer

const (
	Success = 0
	Failure = -1
	Error   = 50000

	ParamErr = 30001

	AccErr       = 40000
	AccAuthErr   = 40001
	AccCreateErr = 40002
	AccSilence   = 40009
	AccNotVerify = 41009

	SignErr = 60001

	WordsPostErr = 41001

	DatabaseErr      = 50002
	CodeEncryptError = 50006
)

var ResMsgFlags = map[int]string{
	Success:      "Success",
	Error:        "Fail",
	ParamErr:     "参数错误",
	WordsPostErr: "发布失败",
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Err  string      `json:"err"`
}

// BuildResponse 响应信息构建
func BuildResponse(code int, data interface{}, msg string, err error) Response {
	var errMsg string
	if err == nil {
		errMsg = ""
	} else {
		errMsg = err.Error()
	}
	return Response{
		Code: code,
		Data: data,
		Msg:  msg,
		Err:  errMsg,
	}
}

// Err 通用错误信息
func Err(code int, msg string) Response {
	return BuildResponse(code, nil, msg, nil)
}

// ParmErr 参数错误信息
func ParmErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return BuildResponse(ParamErr, nil, msg, err)
}

// GetResMsg 获取错误码对应错误信息
func GetResMsg(code int) string {
	msg, ok := ResMsgFlags[code]
	if ok {
		return msg
	}
	return ResMsgFlags[Error]
}
