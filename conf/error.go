package config

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_AUTH         = 20004
	AccountAuthErr     = 40001
	AccountUsernameErr = 40002
)

var ResMsgFlags = map[int]string{
	SUCCESS:    "ok",
	ERROR:      "fail",
	ERROR_AUTH: "Login Error",
}

func GetResMsg(code int) string {
	msg, ok := ResMsgFlags[code]
	if ok {
		return msg
	}
	return ResMsgFlags[ERROR]
}
