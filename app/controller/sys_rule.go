package controller

import (
	"go-api/app/models"
	"go-api/app/request"
	"go-api/app/services"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

func GetRules(c *gin.Context) {
	maps := make(map[string]interface{})

	ret := services.GetRoleList(maps, tool.DefaultGetOffset(c), 10)
	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	}
	tool.JSONP(c, 0, "查询成功", ret["data"])
}

func CreateRule(c *gin.Context) {

	var rule models.AdminRule
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

func UpdateRule(c *gin.Context) {
	var rule models.AdminRule
	c.ShouldBind(&rule)
	if err := request.Verify(rule, request.RuleUpVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.UpdateRule(rule); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}

func DeleteRule(c *gin.Context) {
	var rule models.AdminRule
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
