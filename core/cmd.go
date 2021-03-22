package core

import (
	"fmt"
	"go-api/global"
	"io/ioutil"
	"os"
	"os/exec"
)

func CmdRun() {
	len := len(os.Args)
	if len < 2 {
		Start()
	} else {
		switch os.Args[1] {
		case "start":
			Start()
		case "daemon":
			Daemon()
		case "restart":
			Rstart()
		case "stop":
			Stop()
		}
	}
}

func Daemon() {

	cmd := exec.Command(os.Args[0])
	cmd.Start()
	fmt.Println("[PID]", cmd.Process.Pid)
	//C:/Users/baibai/go/www/src/Go-Gin-Api/app.bsv
	/* cmd := exec.Command("cmd", "/C", "app.bsv")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s failed with error:%s", cmd, err.Error())
		return
	}
	fmt.Printf("%s finished with output:\n%s", cmd, string(output)) */
}

func Start() {
	if pid := ReadPid(global.CF.App.PidPath); pid != "" {
		fmt.Println("服务正在运行中 pid：", pid)
	} else {
		fmt.Println("服务启动中")
		RunServer()
	}
}

func Rstart() {
	if pid := ReadPid(global.CF.App.PidPath); pid != "" {
		fmt.Println("服务平滑重启")
		Restart()
	} else {
		fmt.Println("服务正在启动")
		RunServer()
	}
}

func Stop() {
	if pid := ReadPid(global.CF.App.PidPath); pid != "" {
		fmt.Println("服务关闭 pid：", pid)
		Shutdown()
	} else {
		fmt.Println("服务未启动不需要关闭")
	}
}

func CreatePid(fileName string, pid string) {
	fmt.Println("生成lock文件前")
	err := ioutil.WriteFile(fileName, []byte(pid), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadPid(fileName string) (pid string) {
	body, _ := ioutil.ReadFile(fileName)
	return string(body)
}

func DelPid(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	}
}
