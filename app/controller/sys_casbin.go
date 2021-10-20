package controller

import (
	"fmt"
	"go-api/app/request"
	"go-api/app/services"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

// @Summary 角色分配权限视图
// @Description  角色分配权限视图
// @Router /admin/ruleRbacViwe/:id [get]
func RuleRbacView(c *gin.Context) {
	id := c.Param("id")
	if err := request.VerifyMap(tool.M{"ID": id}, request.CasbinViewVerify); err != nil {
		tool.ViewErr(c, 400, err.Error())
		return
	}

	info := services.GetRuleInfo(id, "ID", "Role_name")
	list := services.GetRuleAuthorityMenuList(id)
	if !info.GetStatus() {
		tool.ViewErr(c, 400, info.GetMsg())
		return
	}
	ret := tool.M{
		"editUrl": "/admin/updateRbac",
		"rule":    info["data"],
		"list":    list["data"],
	}
	tool.HTML(c, "rule/rbac.html", ret)
}

// @Summary 角色分配权限
// @Description  角色分配权限
// @Router /admin/updateRbac  [post]
func UpdateRuleRbac(c *gin.Context) {
	var Casbins request.Casbins
	c.ShouldBind(&Casbins)
	if err := request.Verify(Casbins, request.CasbinUpVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), Casbins)
		return
	}

	var CasbinInfos []request.CasbinInfo
	for _, v := range Casbins.CasbinInfos {
		ruleinfo := services.GetMenuInfo(tool.IntToString(v), "path")
		if !ruleinfo.GetStatus() {
			tool.JSONP(c, 40001, "非法权限菜单", ruleinfo)
			return
		}

		data := tool.StructToMap(ruleinfo["data"])
		CasbinInfos = append(CasbinInfos, request.CasbinInfo{Path: data["Path"].(string), Method: "get|post"})
	}

	//去分配权限
	if ret := services.UpdatePolicyByRuleId(Casbins.RuleId, CasbinInfos); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}

// @Summary 用户角色列表视图
// @Description  用户角色列表视图
// @Router /admin/userRuleView/:id  [get]
func GetUserRules(c *gin.Context) {
	id := c.Param("id")
	if err := request.VerifyMap(tool.M{"ID": id}, request.CasbinViewVerify); err != nil {
		tool.ViewErr(c, 400, err.Error())
		return
	}

	info := services.GetUserByID(id)
	list := services.GetRuleByUserid(id)
	if !info.GetStatus() {
		tool.ViewErr(c, 400, info.GetMsg())
		return
	}

	fmt.Println(list)
	data := tool.M{
		"editUrl": "/admin/setUserRules",
		"info":    info["data"],
		"list":    list["data"],
	}

	tool.HTML(c, "user/user_rule_list.html", data)
}

// @Summary 用户角色列表视图
// @Description  用户角色列表视图
// @Router /admin/setUserRules  [post]
func DistributionUserRules(c *gin.Context) {
	var Casbins request.CasbinsRules
	c.ShouldBind(&Casbins)
	if err := request.Verify(Casbins, request.UserUpCasbinVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), Casbins)
		return
	}

	//去分配权限
	if ret := services.UpRulesForUser(Casbins.UserId, Casbins.Rules); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}
