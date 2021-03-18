package core

import (
	"fmt"
	"go-api/global"
	"os"
)

func CmdRun() {
	len := len(os.Args)
	if len < 2 {
		RunServer()
	} else {
		switch os.Args[1] {
		case "start":
			if global.VP.GetInt("Pid") > 0 {
				fmt.Println("服务正在运行中")
			} else {
				fmt.Println("服务启动中")
				RunServer()
			}
		case "restart":
			if global.VP.GetInt("Pid") > 0 {
				fmt.Println("服务平滑重启")
			} else {
				fmt.Println("服务未启动")
				fmt.Println("服务直接启动")
				RunServer()
			}
		case "stop":
			if global.VP.GetInt("Pid") > 0 {
				fmt.Println("服务关闭")
			} else {
				fmt.Println("服务未启动不需要关闭")
			}
		}
	}
	fmt.Println(os.Args[0])
}
