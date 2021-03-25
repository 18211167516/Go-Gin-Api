package tool

import "fmt"

const (
	// SUCCESS code
	SUCCESS = 0
	// ERROR code
	ERROR = 500
	// INVALID_PARAMS code
	INVALID_PARAMS = 400
	// CUSTOM_ERROR code
	CUSTOM_ERROR = 40001

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_AUTH_CHECK_TOKEN_EMPTY   = 20005
)

var MsgFlags = map[int]string{
	SUCCESS:                        "%s",
	ERROR:                          "%s",
	INVALID_PARAMS:                 "参数%s",
	CUSTOM_ERROR:                   "%s",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_AUTH_CHECK_TOKEN_EMPTY:   "Token参数不能为空",
}

func GetMsg(code int, Msg ...interface{}) string {
	msg, ok := MsgFlags[code]
	if ok {
		return fmt.Sprintf(msg, Msg...)
	}

	return MsgFlags[ERROR]
}
