package tool

import (
	"github.com/gin-gonic/gin"
)

type M map[string]interface{}

type RetData struct{
	Status bool `json:"status"`
	Msg  string `json:"msg"`
	Data  interface{} `json:"data"`
}

func JSONP(c *gin.Context, code int, msg string, data interface{}) {
	c.JSONP(200, gin.H{"error_code": code, "msg": GetMsg(code, msg), "data": data})
}

func DataReturn(status bool,msg string,data interface{}) M{
	result :=M{
		"status" : status,
		"msg" : msg,
		"data" : data,
	}
	return result
}

func (m M) GetStatus() bool{
	return m["status"].(bool)
}

func (m M) GetMsg() string{
	return m["msg"].(string)
}
