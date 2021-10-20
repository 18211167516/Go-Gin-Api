package initialize

import (
	"fmt"
	"go-api/global"
	inLog "go-api/initialize/internal/log"

	log "github.com/sirupsen/logrus"
)

var (
	// Formatter...
	Formatter = map[string]log.Formatter{
		"json": &log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		"text": &log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		"test": &inLog.TestFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	}
)

func defaultFormatter() {
	SetFormatter(global.CF.Log.Formatter)
}

func defaultLevel() {
	SetLevel(global.CF.Log.Level)
}

func defaultReportCaller() {
	SetReportCaller(global.CF.Log.ReportCaller)
}

func SetReportCaller(b bool) {
	log.SetReportCaller(b)
}

func SetFormatter(formatter string) {
	if Formatter, ok := Formatter[formatter]; !ok {
		panic(fmt.Errorf("Log Formatter %s", formatter))
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

func Logrus() *log.Logger {
	logger := log.StandardLogger()
	defaultReportCaller()
	defaultFormatter()
	defaultLevel()
	log.AddHook(&inLog.TestHook{})
	return logger
}
