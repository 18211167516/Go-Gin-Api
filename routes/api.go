package routes

import (
	"go-api/app/controller"
)

func ApiRoute(){
	apiv1 := R.Group("/api/v1")
	{
		//获取文章列表
		apiv1.GET("/users", controller.GetUsers)
		//获取指定文章
		apiv1.GET("/user/:id", controller.GetUser)
        //新建文章
        apiv1.POST("/users", controller.AddUser)
        //更新指定文章
        apiv1.PUT("/users/:id", controller.EditUser)
        //删除指定文章
		apiv1.DELETE("/users/:id", controller.DeleteUser)
	}
}