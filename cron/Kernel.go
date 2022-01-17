package cron

import (
	"go-api/global"
)

type testJob struct{}

func (t testJob) Run() {
	global.LOG.Info("TEST")
}

func Schedule() {
	//每分钟
	global.CRON.AddFunc("@every 1m", "测试1", func() {
		global.LOG.Info("我是谁")
	})

	//每分钟第15秒
	global.CRON.AddJob("15 * * * * *", "测试2", testJob{})
}
