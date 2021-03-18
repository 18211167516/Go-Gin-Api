package core

import (
	"fmt"
	"go-api/global"
	"go-api/initialize"
	"os"
)

func RunServer() {
	//初始化路由
	s := GetServer()
	global.VP.Set("Pid", os.Getpid())
	fmt.Printf(`
欢迎使用 Go-Gin-Api
当前版本:V0.1.1
Server run :%d
PID :%d
`, global.CF.Server.HttpPort, global.VP.Get("Pid"))
	global.LOG.Error(s.ListenAndServe().Error())
	//写入lock文件
}

func GetServer() global.Server {
	if global.SER == nil {
		r := initialize.Routers()
		global.SER = initServer(fmt.Sprintf("%s:%d", global.CF.Server.HttpAddress, global.CF.Server.HttpPort), r)
	}
	return global.SER
}

func Restart() {
	s := GetServer()
	if err := s.Restart(); err != nil {
		global.LOG.Error(err)
	}
}

func Shutdown() {
	s := GetServer()
	s.Shutdown()
}
