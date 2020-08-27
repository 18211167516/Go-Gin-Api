package tool

import (
	"fmt"

	"go-api/config"
	inLog "go-api/tool/internal/log"

	log "github.com/sirupsen/logrus"
)

var (
	// Formatter...
	Formatter = map[string]log.Formatter{
		"json": &log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		"text": &log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"},
		"test": &inLog.TestFormatter{TimestampFormat: "2006-01-02 15:04:05"},
	}

	Log       = log.StandardLogger()
	ApiLog    = apiLog()
	MysqlLog  = mysqlLog()
	AccessLog = accessLog()
)

type F = log.Fields

func defaultFormatter() {
	SetFormatter(config.LogSetting.Formatter)
}

func defaultLevel() {
	SetLevel(config.LogSetting.Level)
}

func defaultReportCaller() {
	SetReportCaller(config.LogSetting.ReportCaller)
}

func SetReportCaller(b bool) {
	log.SetReportCaller(b)
}

func SetFormatter(formatter string) {
	if Formatter, ok := Formatter[formatter]; !ok {
		panic(fmt.Errorf("Log Formatter %s", "unkonm"))
	} else {
		log.SetFormatter(Formatter)
	}
}

func SetLevel(level string) {
	if level, err := log.ParseLevel(level); err != nil {
		panic(fmt.Errorf("Log Level %s", err))
	} else {
		log.SetLevel(level)
	}
}

func init() {
	defaultReportCaller()
	defaultFormatter()
	defaultLevel()
	//log.AddHook(&inLog.TestHook{})
}

func apiLog() *log.Entry {
	return log.WithFields(F{"topic": "api"})
}

func mysqlLog() *log.Entry {
	return log.WithFields(F{"topic": "mysql"})
}
func accessLog() *log.Entry {
	return log.WithFields(F{"topic": "access"})
}
