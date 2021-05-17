package utils

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
	Success:      "ok",
	Error:        "fail",
	WordsPostErr: "发布失败",
}

func GetResMsg(code int) string {
	msg, ok := ResMsgFlags[code]
	if ok {
		return msg
	}
	return ResMsgFlags[Error]
}
