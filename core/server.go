package core

import (
	"fmt"
	"go-api/global"
	"go-api/initialize"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
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
	go CreatePid(global.CF.App.PidPath, strconv.Itoa(os.Getpid()))
	go handleSignals()
	defer DelPid(global.CF.App.PidPath)
	global.LOG.Error(s.ListenAndServe().Error())
	//写入lock文件
}

//解决非后台运行启动时 window下 ctrl c关闭信号
func handleSignals() {
	sysType := runtime.GOOS
	if sysType == "windows" {
		// windows系统
		var sig os.Signal
		signalChan := make(chan os.Signal)
		signal.Notify(
			signalChan,
			syscall.SIGINT,
			syscall.SIGTERM,
		)

		for {
			sig = <-signalChan
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT:
				Shutdown()
			default:
			}
		}
	}
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
	global.LOG.Infoln("正在关闭服务")
	s := GetServer()
	DelPid(global.CF.App.PidPath)
	//删除lock文件
	s.Shutdown()
	os.Exit(1)
}
