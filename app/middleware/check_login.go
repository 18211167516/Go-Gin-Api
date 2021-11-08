package middleware

import (
	"go-api/app/response"
	"go-api/global"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if uid, err := tool.NewSecureCookie(c).GetCookie("uid"); uid != "" {
			waitUse := &response.SysLoginUserResponse{}
			tool.JsonToStruct([]byte(uid), &waitUse)
			c.Set("waitUse", waitUse)
			c.Set("uid", waitUse.ID)
			c.Set("uType", waitUse.Type)
		} else {
			global.LOG.Error(err.Error())
			tool.Output(c, 401, "用户未登录", nil)
			c.Abort()
		}
		c.Next()
	}
}
