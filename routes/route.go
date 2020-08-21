package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api/config"
)

//init router
func InitRouter() *gin.Engine{
    r := initGin()
	loadRoute(r)
	return r;
}

// init Gin
func initGin() *gin.Engine{
	//设置gin模式
	gin.SetMode(config.RunMode)
    return gin.Default()
}

// 加载路由
func loadRoute(r *gin.Engine){
	testRoute(r)
	apiRoute(r)
	swagRoute(r)
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
