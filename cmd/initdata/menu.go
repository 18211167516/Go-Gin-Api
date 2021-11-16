package initdata

import (
	"go-api/app/models"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

var astring = "0"

var menu = []models.SysMenu{
	{Model: models.Model{ID: 14, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "菜单管理", Path: "/admin/menusView", ParentId: 29, Hidden: &astring, Sort: 8, Is_view: "1"},
	{Model: models.Model{ID: 15, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "新增基础菜单", Path: "/admin/createBaseMenu", ParentId: 14, Hidden: &astring, Sort: 8, Is_view: "0"},
	{Model: models.Model{ID: 16, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "角色管理", Path: "/admin/rulesView", ParentId: 29, Hidden: &astring, Sort: 2, Is_view: "1"},
	{Model: models.Model{ID: 18, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "角色列表信息", Path: "/admin/getRules", ParentId: 16, Hidden: &astring, Sort: 3, Is_view: "0"},
	{Model: models.Model{ID: 20, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "更新角色", Path: "/admin/updateRule/:id", ParentId: 16, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 21, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "创建子菜单", Path: "/admin/createChildMenu/:parent_id", ParentId: 14, Hidden: &astring, Sort: 2, Is_view: "0"},
	{Model: models.Model{ID: 24, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "更新菜单", Path: "/admin/updateMenu/:id", ParentId: 14, Hidden: &astring, Sort: 8, Is_view: "0"},
	{Model: models.Model{ID: 25, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "删除菜单", Path: "/deleteMenu/:id", ParentId: 14, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 26, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "创建角色", Path: "/admin/createRule", ParentId: 16, Hidden: &astring, Sort: 2, Is_view: "0"},
	{Model: models.Model{ID: 27, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "删除角色", Path: "/admin/deleteRule/:id", ParentId: 16, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 28, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "权限管理", Path: "/admin/ruleRbacViwe/:id", ParentId: 16, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 29, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "系统管理", Path: "javascript:void(0)", ParentId: 0, Hidden: &astring, Sort: 0, Is_view: "1"},
	{Model: models.Model{ID: 33, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "用户管理", Path: "/admin/usersView", ParentId: 29, Hidden: &astring, Sort: 0, Is_view: "1"},
	{Model: models.Model{ID: 34, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "用户列表信息", Path: "/admin/getUsers", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 35, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "设置角色权限", Path: "/admin/updateRbac", ParentId: 16, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 36, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "新建用户", Path: "/admin/createUser", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 37, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "删除用户", Path: "/admin/deleteUser/:id", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 38, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "编辑用户", Path: "/admin/updateUser/:id", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 39, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "重置管理员密码", Path: "/admin/changePassword/:id", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 41, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "用户配置角色页", Path: "/admin/userRuleView/:id", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 42, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "分配角色", Path: "/admin/setUserRules", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 43, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "个人信息页", Path: "/admin/userView", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 44, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "更新个人信息", Path: "/admin/changeOwnInfo", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 45, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "重置个人密码", Path: "/admin/changeOwnPassword/:id", ParentId: 33, Hidden: &astring, Sort: 0, Is_view: "0"},
	{Model: models.Model{ID: 48, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "后台首页", Path: "/admin/main", ParentId: 0, Hidden: &astring, Sort: 10, Is_view: "1"},
}

func InitSysMenus(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{29, 48}).Find(&[]models.SysMenu{}).RowsAffected == 2 {
			log.Println("sys_menus表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menu).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		log.Printf("[Mysql]--> sys_menus 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
