package main

import (
	"embed"
	"go-api/core"
	"go-api/global"
	"go-api/initialize"
)

//go:embed static templates
var f embed.FS

// @title go-api 框架
// @version 1.0
// @description gin-web框架
// @termsofservice https://github.com/18211167516/Go-Gin-Api
// @contact.name baichonghua
// @contact.email 18211167516@163.com
// @host 127.0.0.1:8080

func main() {
	//初始化FS
	global.FS = f
	//初始化配置
	global.VP = core.Viper()
	//初始化日志
	global.LOG = initialize.Logrus()
	//初始化DB
	global.DB = initialize.Gorm()
	//主进程结束前关闭数据库链接
	sqlDB, _ := global.DB.DB()
	defer sqlDB.Close()
	core.CmdRun()
	//启动服务器
	//core.RunServer()
}
