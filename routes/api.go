package routes

import (
	"go-api/app/controller"
	"go-api/app/middleware"
	"go-api/global"
	"go-api/tool"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func apiRoute(r *gin.Engine) {
	r.GET("/authToken", func(c *gin.Context) {
		p, _ := time.ParseDuration(global.CF.App.JwtExpiresAt)
		expireTime := time.Now().Add(p).Unix()

		claims := tool.Claims{
			ID:       1,
			Username: "白葱花",
			RuleID:   "888",
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

	apiv1 := r.Group("/api/v1", middleware.ApiLogger(), middleware.JWT(), middleware.Casbin_rbac())
	{
		//获取用户列表
		apiv1.GET("/users", controller.GetUsers)
		//获取指定用户
		apiv1.GET("/user/:id", controller.GetUser)
		//新增用户
		apiv1.POST("/users", controller.AddUser)
		//更新指定用户
		apiv1.PUT("/users/:id", controller.EditUser)
		//删除指定用户
		apiv1.DELETE("/users/:id", controller.DeleteUser)

		//获取角色列表
		apiv1.POST("/getRules", controller.GetRules)
		//创建角色
		apiv1.POST("/createRule", controller.CreateRule)
		//更新指定角色
		apiv1.POST("/updateRule", controller.UpdateRule)
		//获取角色信息
		apiv1.GET("/getRule/:id", controller.GetUser)
		//删除指定角色
		apiv1.DELETE("/deleteRule/:id", controller.DeleteRule)
		//设置角色权限
		apiv1.POST("/setRuleAuthority", controller.AddUser)
	}

}
