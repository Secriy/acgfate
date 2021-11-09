package serializer

// 用户端错误
const (
	Success errCode = 0  // 成功代码
	Failure errCode = -1 // 通用错误代码

	ParamErr  errCode = 40001 // 参数错误
	FormatErr errCode = 40002 // 格式不正确

	// 4002x 账号注册时错误

	RegNameExist        errCode = 40021 // 用户名被占用
	EmailExist          errCode = 40022 // 邮箱被占用
	VerifyCodeExpired   errCode = 40023 // 邮箱验证码已过期
	VerifyCodeIncorrect errCode = 40024 // 验证码不正确
	VerifyAlready       errCode = 40025 // 邮箱已验证
	AccNotLegalErr      errCode = 40030 // 账号非法操作
	AccAuthErr          errCode = 40031 // 账号未登录
	AccNotVerify        errCode = 40032 // 账号邮箱未验证
	AccBanErr           errCode = 40032 // 账号被封禁
	AccSilenceErr       errCode = 40033 // 账号被禁言

	WordsPostErr errCode = 41001

	DatabaseErr      errCode = 50002
	CodeEncryptError errCode = 50006
)

// 服务端错误
const (
	Error errCode = 50000

	QueryDBErr       errCode = 50001 // 查询数据失败
	InsertDBErr      errCode = 50002 // 插入数据失败
	UpdateDBErr      errCode = 50003 // 更新数据失败
	DeleteDBErr      errCode = 50004 // 删除数据失败
	CacheStoreErr    errCode = 50011 // 缓存失败
	PasswdEncryptErr errCode = 50021 // 密码加密失败
	TokenGenerateErr errCode = 50022 // Token生成失败
	MailSendErr      errCode = 50031 // 邮件发送失败
)

var ResMsgFlags = map[errCode]string{
	Success:             "成功",
	Failure:             "失败",
	Error:               "出错",
	VerifyCodeExpired:   "未发送验证码或验证码已过期",
	VerifyCodeIncorrect: "验证码不正确",
	VerifyAlready:       "邮箱已验证",
	AccNotVerify:        "邮箱未验证",
	AccAuthErr:          "账号未登录",
	ParamErr:            "参数错误",
	FormatErr:           "格式不正确",
	RegNameExist:        "用户名已存在",
	EmailExist:          "邮箱已存在",
	WordsPostErr:        "发布失败",
	QueryDBErr:          "查询数据失败",
	InsertDBErr:         "插入数据失败",
	UpdateDBErr:         "更新数据失败",
	DeleteDBErr:         "删除数据失败",
	CacheStoreErr:       "缓存存储失败",
	PasswdEncryptErr:    "密码加密失败",
	TokenGenerateErr:    "Token生成失败",
	MailSendErr:         "邮件发送失败",
}

type errCode int32

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

// Message 获取错误码对应错误信息
func Message(code errCode) string {
	msg, ok := ResMsgFlags[code]
	if ok {
		return msg
	}
	return ResMsgFlags[Error]
}

// SuccessResponse 通用成功信息
func SuccessResponse() Response {
	return BuildResponse(Success, nil, Message(Success))
}

// MsgResponse 自定义消息错误返回
func MsgResponse(code interface{}, msg string) Response {
	return BuildResponse(code, nil, msg)
}

// ErrResponse 通用错误信息
func ErrResponse(errCode errCode) Response {
	msg := Message(errCode)
	return BuildResponse(errCode, nil, msg)
}
