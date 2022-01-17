package initialize

import (
	cron "github.com/18211167516/robfig-cron/v3"
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.SugaredLogger
}

func (l Logger) Print(args ...interface{}) {
	l.log.Info(args...)
}

func Cron(log *zap.SugaredLogger) *cron.Cron {
	newlog := Logger{
		log: log,
	}
	return cron.New(cron.WithChain(cron.Recover(newlog), cron.DelayIfStillRunning(newlog), cron.SkipIfStillRunning(newlog)), cron.WithSeconds(), cron.WithLogger(newlog))
}
