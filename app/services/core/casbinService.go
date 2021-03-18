package core

import (
	"go-api/app/request"
	"go-api/global"
	"go-api/tool"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

/*获取某个角色的权限*/
func GetPolicyPathByAuthorityId(ruleID string) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, ruleID)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

/*更新某个角色的权限*/
func UpdatePolicyByRuleId(ruleID string, casbinInfos []request.CasbinInfo) tool.M {
	ClearPolicy(0, ruleID)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{ruleID, v.Path, v.Method})
	}
	e := Casbin()
	success, _ := e.AddPolicies(rules)
	if success {
		return tool.DataReturn(true, "查询成功", nil)
	} else {
		return tool.DataReturn(false, "存在相同api,添加失败,请联系管理员", nil)
	}
}

/*删除Policy*/
func ClearPolicy(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

/*获取Casbin*/
func Casbin() *casbin.Enforcer {
	admin := global.CF.Mysql
	a, _ := gormadapter.NewAdapter("mysql", admin.MysqlUser+":"+admin.MysqlPassword+"@("+admin.MysqlHost+")/"+admin.MysqlName, true)
	e, _ := casbin.NewEnforcer(global.CF.Casbin.ModelPath, a)
	e.LoadPolicy()
	return e
}
