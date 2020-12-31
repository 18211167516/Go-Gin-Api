package core

import (
	"fmt"
	"go-api/global"
	"go-api/initialize"
	"log"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	//初始化路由
	r := initialize.Routers()
	s := initServer(fmt.Sprintf("%s:%d", global.CF.Server.HttpAddress, global.CF.Server.HttpPort), r)
	err := s.ListenAndServe()
	log.Printf("Listen:%s:%s %s\n", global.CF.Server.HttpAddress, global.CF.Server.HttpPort, err)
}
