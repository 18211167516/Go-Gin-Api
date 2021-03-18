package middleware

import (
	"go-api/app/services/core"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

func Casbin_rbac() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*tool.Claims)

		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.GetRuleID()
		// 获取casbin
		casbin := core.Casbin()

		ok, _ := casbin.Enforce(sub, obj, act)
		if !ok {
			c.JSON(200, gin.H{
				"code": tool.ERROR,
				"msg":  "权限不足",
				"data": []string{sub, obj, act},
			})

			c.Abort()
			return
		}

		c.Next()

	}
}
