package initialize

import (
	"go-api/global"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Xorm() *xorm.EngineGroup {
	return XormMysql()
}

func XormMysql() *xorm.EngineGroup {
	dsn := getMasterDsn()
	conns := []string{
		dsn, // 第一个默认是master
		dsn,
	}
	engine, err := xorm.NewEngineGroup("mysql", conns, xorm.LeastConnPolicy())
	if err != nil {
		global.LOG.Error("Mysql启动异常", dsn, err)
		os.Exit(0)
		return nil
	} else {
		engine.SetMaxIdleConns(global.VP.GetInt("mysql.global.MaxIdleConns"))
		engine.SetMaxOpenConns(global.VP.GetInt("mysql.global.MaxOpenConns"))
		engine.SetConnMaxLifetime(global.VP.GetDuration("mysql.global.MaxLifetime"))

		engine.ShowSQL(global.VP.GetBool("mysql.global.LogMode"))
		engine.ShowExecTime(global.VP.GetBool("mysql.global.LogMode"))
		logger := xorm.NewSimpleLogger(global.LOG.Out)
		engine.SetLogger(logger)
	}

	return engine

}
