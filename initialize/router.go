package initialize

import (
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	return routes.InitRouter()
}
