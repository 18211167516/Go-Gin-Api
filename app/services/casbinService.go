package services

import (
	"fmt"
	"go-api/app/request"
	"go-api/global"
	"go-api/tool"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once
var cas *casbin.Enforcer

/*获取某个角色的权限*/
func GetPolicyPathByRuleId(ruleID string) (pathMaps []request.CasbinInfo) {
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

//更新用户角色
func UpRulesForUser(user string, roles []string) tool.M {
	ClearRules(user)
	e := Casbin()
	success, _ := e.AddRolesForUser(user, roles)
	if success {

		return tool.DataReturn(true, "配置角色成功", nil)
	} else {
		return tool.DataReturn(false, "配置角色失败", nil)
	}
}

//确定用户是否具有角色
func HasRuleForUser(user string, rule string) bool {
	e := Casbin()
	success, _ := e.HasRoleForUser(user, rule)
	return success
}

//获取用户所有角色
func GetRulesForUser(user string) (list []string) {
	e := Casbin()
	list, _ = e.GetRolesForUser(user)
	return list
}

func HasPolicyByRuleIdPath(ruleID string, path string, method string) bool {
	e := Casbin()
	return e.HasNamedPolicy("p", ruleID, path, method)
}

func HasPolicyByRuleIdsPath(RuleID []string, path string, method string) bool {
	e := Casbin()

	for _, v := range RuleID {
		if e.HasNamedPolicy("p", v, path, method) {
			return true
		}
	}
	return false
}

/*更新某个角色的权限*/
func UpdatePolicyByRuleId(ruleID string, casbinInfos []request.CasbinInfo) tool.M {
	ClearPolicy(0, ruleID)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{ruleID, v.Path, v.Method})
	}
	e := Casbin()
	//添加角色权限
	success, _ := e.AddNamedPolicies("p", rules)
	if success {
		return tool.DataReturn(true, "配置权限成功", nil)
	} else {
		return tool.DataReturn(false, "配置权限失败", nil)
	}
}

//删除某个角色
func ClearRule(role string) bool {
	e := Casbin()
	success, _ := e.DeleteRole(role)
	return success
}

//删除某个全部角色
func ClearRules(user string) bool {
	e := Casbin()
	success, _ := e.DeleteRolesForUser(user)
	return success
}

/*删除Policy*/
func ClearPolicy(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

func getDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		global.VP.GetString("mysql.master.User"),
		global.VP.GetString("mysql.master.Password"),
		global.VP.GetString("mysql.master.Host"),
		global.VP.GetString("mysql.master.DBName"),
		global.VP.GetString("mysql.master.Config"),
	)
}

/*获取Casbin*/
func Casbin() *casbin.Enforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapter("mysql", getDsn(), true)
		cas, _ = casbin.NewEnforcer(global.VP.GetString("casbin.ModelPath"), a)
		cas.LoadPolicy()
	})
	return cas
}
