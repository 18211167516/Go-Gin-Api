package routes

import (
	"github.com/gin-gonic/gin"

	"go-api/app/middleware"
	"go-api/global"
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
	gin.SetMode(global.CF.RunMode)
	engine := gin.New()
	engine.Use(middleware.ApiLogger(), middleware.Recovery())
	return engine
}

// 加载路由
func loadRoute(r *gin.Engine) {
	testRoute(r)
	apiRoute(r)
	swagRoute(r)
	fileRoute(r)
}
