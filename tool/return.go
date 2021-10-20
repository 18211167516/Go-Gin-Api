package tool

import (
	"fmt"

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

func HTML(c *gin.Context, name string, data interface{}) {
	c.HTML(200, name, data)
}

func ViewErr(c *gin.Context, code int, message string) {
	HTML(c, "error/404.html", M{"code": code, "message": message})
}

func Output(c *gin.Context, code int, msg string, data interface{}) {
	ContentType := c.ContentType()
	if c.Request.Method == "GET" && (ContentType == "" || ContentType == "application/html") {
		if code == 401 {
			c.Redirect(302, "/admin/login")
		} else {
			c.Redirect(302, fmt.Sprintf("/admin/error/%d/%s", code, msg))
		}
	} else {
		JSONP(c, code, msg, data)
		c.Abort()
	}
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
