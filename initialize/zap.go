package initialize

import (
	"fmt"
	"go-api/global"
	"go-api/tool"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level zapcore.Level

func ZapSugar(logName string) *zap.SugaredLogger {
	return Zap(logName).Sugar().Named("Cron")
}

func Zap(logName string) (logger *zap.Logger) {
	if ok, _ := tool.PathExists(global.CF.Log.LogDir); !ok { // 判断是否有LogDir文件夹
		fmt.Printf("create %v directory\n", global.CF.Log.LogDir)
		_ = os.Mkdir(global.CF.Log.LogDir, os.ModePerm)
	}
	switch global.CF.Log.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if logName == "" {
		logName = "gga"
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(logName), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore(logName))
	}
	if global.CF.Log.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.CF.Log.Formatter == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())

	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(logName string) (core zapcore.Core) {
	writer, err := getWriteSyncer(logName) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02-15:04:05"))
}

//分割日志文件
func getWriteSyncer(logName string) (zapcore.WriteSyncer, error) {
	if global.CF.Log.OutFile {
		fileWriter, err := zaprotatelogs.New(
			path.Join(global.CF.Log.LogDir, "%Y-%m-%d."+logName+".log"),
			zaprotatelogs.WithMaxAge(7*24*time.Hour),
			zaprotatelogs.WithRotationTime(24*time.Hour),
		)
		if global.CF.Log.LogInConsole {
			return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
		}
		return zapcore.AddSync(fileWriter), err
	} else {
		return zapcore.AddSync(os.Stdout), nil
	}

}
