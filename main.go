package main

import (
    "go-api/routes"
    _ "go-api/app/request"
)

func main() {
    //启动服务器
    routes.Run()
}