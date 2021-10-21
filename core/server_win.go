// +build windows

package core

import (
	"context"
	"go-api/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WinServer struct {
	*http.Server
	//shutdownChan chan bool
}

func NewHotServer(server *http.Server) (srv *WinServer) {
	return &WinServer{
		Server: server,
	}
}

func (srv *WinServer) Shutdown() {
	if err := srv.Server.Shutdown(context.Background()); err != nil {
		global.LOG.Error("HTTP server shutdown error: %v", zap.Any("err", err))
	} else {
		global.LOG.Info("HTTP server shutdown success.")
	}
}

// start new process to handle HTTP Connection
func (srv *WinServer) Restart() (err error) {
	global.LOG.Error("windows 下不支持restart")
	return nil
}

func initServer(address string, router *gin.Engine) global.Server {
	http := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return NewHotServer(http)
}
