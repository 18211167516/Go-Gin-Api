package middleware

import (
	"bytes"
	"fmt"
	"go-api/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Option func(*LogParams, *gin.Context)
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

	FormatFunc LogFormatFunc
}

func (Log *LogParams) Run() {
	Log.FormatFunc(Log)
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

type LogFormatFunc func(params *LogParams)

//默认格式化方式
var defaultLogFormatter = func(param *LogParams) {
	Fields := []zap.Field{
		zap.Int("http_code", param.StatusCode),
		zap.String("exec_time", fmt.Sprintf("%13v", param.Latency)),
		zap.String("ip", param.ClientIP),
		zap.String("method", param.Method),
		zap.String("url", param.Path),
	}
	global.LOG.Named("Admin").Info(param.ErrorMessage, Fields...)
}

/*api格式化输出*/
var apiLogFormatter = func(param *LogParams) {
	Fields := []zap.Field{
		zap.Int("http_code", param.StatusCode),
		zap.String("exec_time", fmt.Sprintf("%13v", param.Latency)),
		zap.String("ip", param.ClientIP),
		zap.String("method", param.Method),
		zap.String("url", param.Path),
		zap.String("request", param.request),
		zap.String("Response", param.Response),
	}
	global.LOG.Named("Api").Info(param.ErrorMessage, Fields...)
}

/*默认log兼容*/
func DefaultLog() gin.HandlerFunc {
	opt := func(L *LogParams, c *gin.Context) {
		ContentType := c.ContentType()
		if c.Request.Method == "GET" && (ContentType == "" || ContentType == "application/html") {
			L.FormatFunc = defaultLogFormatter
		} else {
			L.FormatFunc = apiLogFormatter
		}

	}
	return LoggerWith(opt)
}

func Logger() gin.HandlerFunc {
	return LoggerWithFormatterLogFunc(defaultLogFormatter)
}

func ApiLogger() gin.HandlerFunc {
	return LoggerWithFormatterLogFunc(apiLogFormatter)
}

/**/
func LoggerWithFormatterLogFunc(f LogFormatFunc) gin.HandlerFunc {
	opt := func(L *LogParams, c *gin.Context) {
		L.FormatFunc = f
	}
	return LoggerWith(opt)
}

/*日志记录*/
func LoggerWith(opts ...Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		logWirter := &LogWriter{ResponseWriter: c.Writer, NewWirter: bytes.NewBufferString("")}
		c.Writer = logWirter
		start := time.Now()
		path := c.Request.URL.String()
		// Process request
		c.Next()

		param := &LogParams{
			Request: c.Request,
		}
		for _, opt := range opts {
			opt(param, c)
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

		param.Run()
	}
}
