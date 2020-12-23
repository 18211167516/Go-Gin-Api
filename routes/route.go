package routes

import (
	"github.com/gin-gonic/gin"

	"go-api/app/middleware"
	"go-api/config"
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
