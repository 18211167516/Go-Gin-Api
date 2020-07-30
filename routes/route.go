package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api/config"
)

var R *gin.Engine;//定义全局变量

//init router
func InitRouter() *gin.Engine{
    initGin()
	loadRoute()
	return R;
}

// init Gin
func initGin(){
	//设置gin模式
	gin.SetMode(config.RunMode)
    R = gin.Default()
}

// 加载路由
func loadRoute(){
	TestRoute()
	ApiRoute()
}

//启动服务器
func Run(){
	r := InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d",config.ServerSetting.HttpAddress,config.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    config.ServerSetting.ReadTimeout,//请求响应的超市时间
		WriteTimeout:   config.ServerSetting.WriteTimeout,//返回响应的超时时间
		//MaxHeaderBytes: 1 << 20,//默认的1MB
	}
	s.ListenAndServe()
}
