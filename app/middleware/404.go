package middleware

import (
	"github.com/gin-gonic/gin"

	"go-api/tool"
)

func NoRoute() gin.HandlerFunc {

	return func(c *gin.Context) {
		tool.Output(c, 404, "路由不存在", nil)
		/* ContentType := c.ContentType()
		if c.Request.Method == "GET" && (ContentType == "" || ContentType == "application/html") {
			tool.HTML(c, "error/404.html", tool.M{"code": 404, "message": "很抱歉，但是那个页面看起来已经不存在了。"})
		} else {
			tool.JSONP(c, 404, "路由不存在", nil)
			c.Abort()
		} */
	}
}
