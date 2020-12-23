package core

import (
	"fmt"
	"go-api/config"
	"go-api/routes"
	"log"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	r := routes.InitRouter()
	s := initServer(fmt.Sprintf("%s:%d", config.ServerSetting.HttpAddress, config.ServerSetting.HttpPort), r)
	err := s.ListenAndServe()
	log.Printf("Listen: %s\n", err)
}
