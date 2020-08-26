package tests

import (
	. "go-api/tool"
	"testing"
)

func TestLogFormattest(t *testing.T) {
	SetFormatter("test")
	Log.Info("info msg")
}

func TestLogFormattext(t *testing.T) {
	SetFormatter("text")
	Log.Info("info msg")
}

func TestLogFormatjson(t *testing.T) {
	SetFormatter("json")
	Log.Info("info msg")
}

func TestLogHooktest(t *testing.T) {
	SetFormatter("json")
	Log.Info("info msg")
	Log.Debug("debug msg")
	Log.Error("error msg")
	Log.Panic("panic msg")
	Log.Fatal("Fatal msg")
}

func TestLogApi(t *testing.T) {
	SetFormatter("json")
	ApiLog.Info("info msg")
	ApiLog.Debug("debug msg")
	ApiLog.Error("error msg")
	ApiLog.Panic("panic msg")
	ApiLog.Fatal("Fatal msg")
}

func TestLogMysql(t *testing.T) {
	SetFormatter("json")
	MysqlLog.Info("info msg")
	MysqlLog.Debug("debug msg")
	MysqlLog.Error("error msg")
	MysqlLog.Panic("panic msg")
	MysqlLog.Fatal("Fatal msg")
}

func TestLogAccess(t *testing.T) {
	SetFormatter("json")
	AccessLog.Info("info msg")
	AccessLog.Debug("debug msg")
	AccessLog.Error("error msg")
	AccessLog.Panic("panic msg")
	AccessLog.Fatal("Fatal msg")
}
