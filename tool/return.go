package tool

import (
	"github.com/gin-gonic/gin"
)

type M map[string]interface{}

type RetData struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type JSONRET struct {
	Error_code int         `json:"error_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func JSONP(c *gin.Context, code int, msg string, data interface{}) {
	h := JSONRET{
		code, GetMsg(code, msg), data,
	}
	c.JSONP(200, h)
}

func DataReturn(status bool, msg string, data interface{}) M {
	result := M{
		"status": status,
		"msg":    msg,
		"data":   data,
	}
	return result
}

func (m M) GetStatus() bool {
	return m["status"].(bool)
}

func (m M) GetMsg() string {
	return m["msg"].(string)
}
