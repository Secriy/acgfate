package serializer

// An errCode is a 32-bit error code.
type errCode int32

// 用户端错误
const (
	CodeSuccess             errCode = 0     // 成功代码
	CodeFailure             errCode = -1    // 通用错误代码
	CodeParamErr            errCode = 40001 // 参数错误
	CodeFormatErr           errCode = 40002 // 格式不正确
	CodeLoginIncorrect      errCode = 40011 // 账号或密码错误
	CodeRegNameExist        errCode = 40021 // 用户名被占用
	CodeEmailExist          errCode = 40022 // 邮箱被占用
	CodeVerifyCodeExpired   errCode = 40023 // 邮箱验证码已过期
	CodeVerifyCodeIncorrect errCode = 40024 // 验证码不正确
	CodeVerifyAlready       errCode = 40025 // 邮箱已验证
	CodeAccNotLegalErr      errCode = 40030 // 账号非法操作
	CodeAccAuthErr          errCode = 40031 // 账号未登录
	CodeAccNotVerify        errCode = 40032 // 账号邮箱未验证
	CodeAccBanErr           errCode = 40033 // 账号被封禁
	CodeAccSilenceErr       errCode = 40034 // 账号被禁言
)

// 服务端错误
const (
	CodeError            errCode = 50000
	CodeQueryDBErr       errCode = 50001 // 查询数据失败
	CodeInsertDBErr      errCode = 50002 // 插入数据失败
	CodeUpdateDBErr      errCode = 50003 // 更新数据失败
	CodeDeleteDBErr      errCode = 50004 // 删除数据失败
	CodeCacheStoreErr    errCode = 50011 // 缓存失败
	CodePasswdEncryptErr errCode = 50021 // 密码加密失败
	CodeTokenGenerateErr errCode = 50022 // Token生成失败
	CodeMailSendErr      errCode = 50031 // 邮件发送失败
)

var codeFlags = map[errCode]string{
	CodeSuccess:             "成功",
	CodeFailure:             "失败",
	CodeError:               "出错",
	CodeLoginIncorrect:      "账号或密码错误",
	CodeVerifyCodeExpired:   "未发送验证码或验证码已过期",
	CodeVerifyCodeIncorrect: "验证码不正确",
	CodeVerifyAlready:       "邮箱已验证",
	CodeAccNotLegalErr:      "账号非法操作",
	CodeAccNotVerify:        "邮箱未验证",
	CodeAccAuthErr:          "账号未登录",
	CodeAccBanErr:           "账号被封禁",
	CodeAccSilenceErr:       "账号被禁言",
	CodeParamErr:            "参数错误",
	CodeFormatErr:           "格式不正确",
	CodeRegNameExist:        "用户名已存在",
	CodeEmailExist:          "邮箱已存在",
	CodeQueryDBErr:          "查询数据失败",
	CodeInsertDBErr:         "插入数据失败",
	CodeUpdateDBErr:         "更新数据失败",
	CodeDeleteDBErr:         "删除数据失败",
	CodeCacheStoreErr:       "缓存存储失败",
	CodePasswdEncryptErr:    "密码加密失败",
	CodeTokenGenerateErr:    "Token生成失败",
	CodeMailSendErr:         "邮件发送失败",
}

// String 获取错误码对应错误信息
func (e errCode) String() string {
	if s, ok := codeFlags[e]; ok {
		return s
	}
	return codeFlags[CodeError]
}
