package middleware

import (
	"go-api/app/response"
	"go-api/core/session"
	"go-api/global"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := session.Default(c)
		uid := s.Get("waitUse")
		//uid, _ := tool.NewSecureCookie(c).GetCookie("uid")
		if uid != nil {
			uid := uid.(string)
			waitUse := &response.SysLoginUserResponse{}
			tool.JsonToStruct([]byte(uid), &waitUse)
			c.Set("waitUse", waitUse)
			c.Set("uid", waitUse.ID)
			c.Set("uType", waitUse.Type)
		} else {
			global.LOG.Error("用户未登录")
			tool.Output(c, 401, "用户未登录", nil)
			c.Abort()
		}
		c.Next()
	}
}
