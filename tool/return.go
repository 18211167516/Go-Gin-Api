package tool

import (
	"github.com/gin-gonic/gin"
)

func JSONP(c *gin.Context, code int, msg string, data interface{}) {
	c.JSONP(200, gin.H{"error_code": code, "msg": GetMsg(code, msg), "data": data})
}
