package middleware

import (
	"fmt"
	"net"
	"net/http/httputil"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

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
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				message := trace(fmt.Sprintf("%s", err))
				//string(debug.Stack())
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.LOG.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				} else {
					global.LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("key string", message),
					)
				}

				if global.VP.GetString("RunMode") == "release" {
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
