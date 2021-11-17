package controller

import (
	"go-api/app/models"
	"go-api/app/request"
	"go-api/app/services"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

// @Summary 角色列表视图
// @Description  角色列表视图
// @Router /admin/rulesView [get]
func RuleListView(c *gin.Context) {
	uid := c.GetString("uid")
	view_route := c.Request.URL.RequestURI()
	data := tool.M{
		"dataUrl":    "/admin/getRules",
		"dataMethod": "POST",
		"addUrl":     services.GetButtonPermission(uid, view_route, "/admin/createRule"),
		"editUrl":    services.GetButtonPermission(uid, view_route, "/admin/updateRule/:id"),
		"delUrl":     services.GetButtonPermission(uid, view_route, "/admin/deleteRule/:id"),
		"rbacUrl":    services.GetButtonPermission(uid, view_route, "/admin/ruleRbacViwe/:id"),
	}
	tool.HTML(c, "rule/rule_list.html", data)
}

// @Summary 角色列表
// @Description  角色列表
// @Router /admin/getRules [post]
func GetRules(c *gin.Context) {
	maps := make(map[string]interface{})
	ret := services.GetRoleList(maps, tool.DefaultGetOffset(c), 10)
	tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
}

// @Summary 新建角色
// @Description  新建角色
// @Router /admin/createRule [post]
func CreateRule(c *gin.Context) {

	var rule models.SysRule
	c.ShouldBind(&rule)
	if err := request.Verify(rule, request.RuleAddVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.CreateRule(rule); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}

// @Summary 更新角色信息
// @Description  更新角色信息
// @Router /admin/updateRule/:id [post]
func UpdateRule(c *gin.Context) {
	var rule models.SysRule

	c.ShouldBind(&rule)

	id := c.Param("id")
	rule.ID = tool.StringToInt(id)

	if err := request.Verify(rule, request.RuleUpVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.UpdateRule(id, rule); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}

// @Summary 删除角色
// @Description  删除角色
// @Router /admin/deleteRule/:id [post]
func DeleteRule(c *gin.Context) {
	var rule models.SysRule
	c.ShouldBindUri(&rule)
	if err := request.Verify(rule, request.RuleDelVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.DelRule(rule); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}
