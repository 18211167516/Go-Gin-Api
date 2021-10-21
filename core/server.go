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
	当前版本:V1.1.1
	Server run :%d
	PID :%d
	`, global.VP.GetInt("server.HttpPort"), global.VP.Get("Pid"))

	go CreatePid(global.VP.GetString("app.PidPath"), strconv.Itoa(os.Getpid()))
	go handleSignals()
	defer DelPid(global.VP.GetString("app.PidPath"))
	global.LOG.Error(s.ListenAndServe().Error())
	//写入lock文件
}

//解决非后台运行启动时 window下 ctrl c关闭信号
func handleSignals() {
	sysType := runtime.GOOS
	if sysType == "windows" {
		// windows系统
		var sig os.Signal
		signalChan := make(chan os.Signal, 3)
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
		global.SER = initServer(fmt.Sprintf("%s:%d", global.VP.GetString("server.HttpAddress"), global.VP.GetInt("server.HttpPort")), r)
	}
	return global.SER
}

func Restart() {
	s := GetServer()
	if err := s.Restart(); err != nil {
		global.LOG.Error(err.Error())
	}
}

func Shutdown() {
	global.LOG.Info("正在关闭服务")
	s := GetServer()
	DelPid(global.VP.GetString("app.PidPath"))
	//删除lock文件
	s.Shutdown()
	os.Exit(1)
}
