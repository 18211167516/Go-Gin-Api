package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"go-api/app/middleware"
	"go-api/config"
	. "go-api/tool"
)

//init router
func InitRouter() *gin.Engine {
	r := initGin()
	loadRoute(r)
	return r
}

// init Gin
func initGin() *gin.Engine {
	//设置gin模式
	gin.SetMode(config.RunMode)
	engine := gin.New()
	engine.Use(middleware.Logger(), middleware.Recovery())
	return engine
}

// 加载路由
func loadRoute(r *gin.Engine) {
	testRoute(r)
	apiRoute(r)
	swagRoute(r)
}

//启动服务器
func Run() {
	r := InitRouter()

	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.ServerSetting.HttpAddress, config.ServerSetting.HttpPort),
		Handler:      r,
		ReadTimeout:  config.ServerSetting.ReadTimeout,  //请求响应的超市时间
		WriteTimeout: config.ServerSetting.WriteTimeout, //返回响应的超时时间
		//MaxHeaderBytes: 1 << 20,//默认的1MB
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			Log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	Log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		Log.Fatal("Server Shutdown:", err)
	}

	Log.Println("Server exiting")
}
