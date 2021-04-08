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
	m := global.CF.Mysql
	dsn := m.MysqlUser + ":" + m.MysqlPassword + "@tcp(" + m.MysqlHost + ")/" + m.MysqlName + "?" + m.Config
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
		engine.SetMaxIdleConns(10)
		engine.SetMaxOpenConns(100)
		engine.SetConnMaxLifetime(m.MaxLifetime)

		engine.ShowSQL(global.CF.Mysql.LogMode)
		engine.ShowExecTime(global.CF.Mysql.LogMode)
		logger := xorm.NewSimpleLogger(global.LOG.Out)
		engine.SetLogger(logger)
	}

	return engine

}
