// +build !windows

package core

import (
	"time"

	"github.com/18211167516/hotstart"
	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	s := hotstart.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
