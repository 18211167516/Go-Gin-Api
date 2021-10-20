package middleware

import (
	"go-api/app/response"
	"go-api/app/services"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

func Casbin_rbac() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("waitUse")
		User := claims.(*response.SysLoginUserResponse)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取casbin
		casbin := services.Casbin()
		ok, err := casbin.Enforce(User, obj, tool.ToLower(act))
		if !ok {
			if err != nil {
				tool.Output(c, 419, err.Error(), nil)
			}
			tool.Output(c, 419, "权限不足", nil)
			return
		}

		c.Next()

	}
}
