package main

import (
    "go-api/config"
    "go-api/routes"
)

func main() {
    //初始化配置
    config.InitConfig()
    //启动服务器
    routes.Run()
}