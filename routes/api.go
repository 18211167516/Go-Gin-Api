package routes

import (
	"go-api/app/controller"
	"go-api/app/middleware"
	"go-api/app/models"
	"go-api/app/services"
	"go-api/global"
	"go-api/tool"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func apiRoute(r *gin.Engine) {
	r.GET("/authToken", middleware.DefaultLog(), middleware.Recovery(), func(c *gin.Context) {
		p, _ := time.ParseDuration(global.CF.App.JwtExpiresAt)
		expireTime := time.Now().Add(p).Unix()

		ret := services.GetUserByID("1")

		if !ret.GetStatus() {
			tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
			return
		}

		waitUse := ret["data"].(models.SysUser)

		claims := tool.Claims{
			ID:       "1",
			Name:     waitUse.Name,
			Type:     waitUse.Type,
			RuleName: "超级管理员",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expireTime, //过期时间
				Issuer:    "go-api",   //签发人
			},
		}
		token, err := tool.GenerateToken(claims)
		if err != nil {
			tool.JSONP(c, tool.ERROR_AUTH_TOKEN, "", nil)
			return
		}
		tool.JSONP(c, 0, "成功", token)
	})

	api := r.Group("api", middleware.DefaultLog(), middleware.Recovery())
	{
		api.GET("/captcha", controller.Captcha)
	}
}
