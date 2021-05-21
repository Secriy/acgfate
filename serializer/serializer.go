package serializer

const (
	Success = 0     // 成功代码
	Failure = -1    // 通用错误代码
	Error   = 50000 // 服务端错误通用代码

	ParamErr       = 40011 // 参数错误
	AccNotLegalErr = 40030 // 账号非法操作
	AccAuthErr     = 40031 // 账号未登录
	AccNotVerify   = 40032 // 账号邮箱未验证
	AccBanErr      = 40032 // 账号被封禁
	AccSilenceErr  = 40033 // 账号被禁言

	WordsPostErr = 41001

	DatabaseErr      = 50002
	CodeEncryptError = 50006
)

var ResMsgFlags = map[int]string{
	Success:      "成功",
	Failure:      "失败",
	ParamErr:     "参数错误",
	WordsPostErr: "发布失败",
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

// ErrResponse 通用错误信息
func ErrResponse(code int, msg string) Response {
	return BuildResponse(code, nil, msg)
}

// ParmErr 参数错误信息
func ParmErr() Response {
	return BuildResponse(ParamErr, nil, "参数错误")
}

// GetResMsg 获取错误码对应错误信息
func GetResMsg(code int) string {
	msg, ok := ResMsgFlags[code]
	if ok {
		return msg
	}
	return ResMsgFlags[Error]
}
