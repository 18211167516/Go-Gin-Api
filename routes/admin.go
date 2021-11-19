package routes

import (
	"go-api/app/controller"
	"go-api/app/middleware"

	"github.com/gin-gonic/gin"
)

func adminRoute(r *gin.Engine) {

	//错误页
	r.GET("/admin/error/:code/:message", controller.Error)
	//登录页
	r.GET("/admin/login", controller.Login)
	//登录
	r.POST("/admin/login", controller.Loginin)
	//退出登录
	r.POST("/admin/loginout", controller.Loginout)
	//首页
	r.GET("/admin/index", middleware.CheckLogin(), controller.Index)

	admin := r.Group("/admin", middleware.DefaultLog(), middleware.Recovery(), middleware.CheckLogin(), middleware.Casbin_rbac())
	{
		//main页
		admin.GET("/main", controller.Main)
		//---用户管理---
		//用户列表视图
		admin.GET("/usersView", controller.UsersView)
		//个人信息页
		admin.GET("/userView", controller.UserView)
		//获取用户列表数据
		admin.POST("/getUsers", controller.GetUsers)
		//新增用户
		admin.POST("/createUser", controller.AddUser)
		//更新指定用户
		admin.POST("/updateUser/:id", controller.EditUser)
		//修改密码
		admin.POST("/changePassword/:id", controller.ChangePassword)
		//修改自己密码
		admin.POST("/changeOwnPassword", controller.ChangeOwnPassword)
		//更新自己用户信息
		admin.POST("/changeOwnInfo", controller.ChangeOwnInfo)
		//删除指定用户
		admin.POST("/deleteUser/:id", controller.DeleteUser)

		//---角色管理---
		//角色列表视图
		admin.GET("/rulesView", controller.RuleListView)
		//角色列表数据
		admin.POST("/getRules", controller.GetRules)
		//创建角色
		admin.POST("/createRule", controller.CreateRule)
		//更新指定角色
		admin.POST("/updateRule/:id", controller.UpdateRule)
		//删除指定角色
		admin.POST("/deleteRule/:id", controller.DeleteRule)
		//角色权限页
		admin.GET("/ruleRbacViwe/:id", controller.RuleRbacView)
		//更新角色权限
		admin.POST("/updateRbac", controller.UpdateRuleRbac)
		//用户角色列表
		admin.GET("/userRuleView/:id", controller.GetUserRules)
		//分配用户角色
		admin.POST("/setUserRules", controller.DistributionUserRules)

		//--菜单管理
		//菜单列表视图
		admin.GET("/menusView", controller.MenusView)
		//创建根菜单
		admin.POST("/createBaseMenu", controller.CreateBaseMenu)
		//创建子菜单
		admin.POST("/createChildMenu/:parent_id", controller.CreateChildMenu)
		//更新菜单
		admin.POST("/updateMenu/:id", controller.UpdateMenu)
		//删除菜单
		admin.POST("/deleteMenu/:id", controller.DeleteMenu)
	}

}
