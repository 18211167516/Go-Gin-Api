package routes

import (
	"go-api/app/controller"

	"github.com/gin-gonic/gin"
)

func apiRoute(r *gin.Engine){
	apiv1 := r.Group("/api/v1")
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
	}
}