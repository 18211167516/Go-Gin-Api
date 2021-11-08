package initdata

import (
	"log"
	"os"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var casbin = []gormadapter.CasbinRule{
	{Ptype: "g", V0: "1", V1: "40"},
	{Ptype: "g", V0: "4", V1: "39"},
	{Ptype: "p", V0: "39", V1: "/admin/changeOwnInfo", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/changeOwnPassword", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/getRules", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/getUsers", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/index", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/main", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/menusView", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/rulesView", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/usersView", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "/admin/userView", V2: "get|post"},
	{Ptype: "p", V0: "39", V1: "javascript:void(0)", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/changeOwnInfo", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/changeOwnPassword", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/changePassword/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/createBaseMenu", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/createChildMenu/:parent_id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/createRule", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/createUser", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/deleteRule/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/deleteUser/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/getUsers", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/index", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/main", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/menusView", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/ruleRbacViwe/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/rulesView", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/setUserRules", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/updateMenu/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/updateRbac", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/updateRule/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/updateUser/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/userRuleView/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/usersView", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/admin/userView", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "/deleteMenu/:id", V2: "get|post"},
	{Ptype: "p", V0: "40", V1: "javascript:void(0)", V2: "get|post"},
}

func InitCasbin(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("ptype = ? AND v0 IN ?", "p", []string{"39", "40"}).Find(&[]gormadapter.CasbinRule{}).RowsAffected == 2 {
			log.Println("casbin_rule表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&casbin).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		log.Printf("[Mysql]--> casbin_rule表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
