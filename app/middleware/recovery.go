package middleware

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"

	"go-api/global"
	"go-api/tool"
)

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[1:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				message := trace(fmt.Sprintf("%s", err))
				//string(debug.Stack())
				global.LOG.Error(message)

				if global.CF.RunMode == "release" {
					message = "Internal Server Error"
				}
				ContentType := c.ContentType()
				if c.Request.Method == "GET" && (ContentType == "" || ContentType == "application/html") {
					tool.HTML(c, "error_404.html", tool.M{"code": 500, "message": message})
				} else {
					tool.JSONP(c, 500, message, nil)
				}
			}
		}()
		c.Next()
	}
}
