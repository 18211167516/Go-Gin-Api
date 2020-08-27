package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	. "go-api/tool"
)

type LogParams struct {
	Request *http.Request
	Start   time.Time
	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
	// Response is HTTP response body
	Response string
	// request
	request string
}

type LogWriter struct {
	gin.ResponseWriter
	NewWirter *bytes.Buffer
}

// 为了实现双写
func (w LogWriter) Write(p []byte) (int, error) {
	if n, err := w.NewWirter.Write(p); err != nil {
		return n, nil
	}
	return w.ResponseWriter.Write(p)
}

type LogFormatter func(params LogParams) string

type LogFunc func(params LogParams, f LogFormatter)

var defaultLogFormatter = func(param LogParams) string {

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("%3d| %13v | %15s | %s | %v  | %s",
		//param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}

var defaultLog = func(params LogParams, f LogFormatter) {
	AccessLog.Info(f(params))
}

var apiLog = func(params LogParams, f LogFormatter) {

	ApiLog.WithFields(F{
		"start_time": params.Start.Format("2006/01/02 - 15:04:05"),
		"exec_time":  fmt.Sprintf("%13v", params.Latency),
		"http_code":  params.StatusCode,
		"ip":         params.ClientIP,
		"method":     params.Method,
		"url":        params.Path,
		"request":    params.request,
		"Response":   params.Response,
	}).Info(f(params))
}

func Logger() gin.HandlerFunc {
	return LoggerWithFormatter(defaultLogFormatter)
}

func ApiLogger() gin.HandlerFunc {
	f := func(param LogParams) string {

		return fmt.Sprintf("%s", param.ErrorMessage)
	}
	return LoggerWithFormatterLogFunc(f, apiLog)
}

func LoggerWithFormatterLogFunc(f LogFormatter, l LogFunc) gin.HandlerFunc {
	return LoggerWith(f, l)
}

func LoggerWithFormatter(f LogFormatter) gin.HandlerFunc {
	return LoggerWith(f, defaultLog)
}

func LoggerWith(f LogFormatter, l LogFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		logWirter := &LogWriter{ResponseWriter: c.Writer, NewWirter: bytes.NewBufferString("")}
		c.Writer = logWirter
		start := time.Now()
		path := c.Request.URL.String()
		// Process request
		c.Next()

		param := LogParams{
			Request: c.Request,
		}
		// Stop timer
		param.Start = start
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.Response = logWirter.NewWirter.String()
		param.request = c.Request.PostForm.Encode()

		param.Path = path
		l(param, f)
	}
}
