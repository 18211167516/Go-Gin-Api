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
	fmt.Println("后台运行：[PID]", cmd.Process.Pid)
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
		KillProcess()
	} else {
		fmt.Println("服务未启动不需要关闭")
	}
}

func KillProcess() error {
	global.LOG.Info("Killing browser process")
	kill := exec.Command("taskkill", "/T", "/F", "/PID", ReadPid(global.CF.App.PidPath))
	err := kill.Run()
	if err != nil {
		global.LOG.Error("Error killing chromium process")
	}
	global.LOG.Info("Browser process was killed")
	DelPid(global.CF.App.PidPath)
	return err
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
