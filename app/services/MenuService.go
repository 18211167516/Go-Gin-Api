package services

import (
	"errors"
	"go-api/app/models"
	"go-api/global"
	"go-api/tool"

	"gorm.io/gorm"
)

//获取基础菜单树 Children
func getBaseMenuTreeMap(where map[string]interface{}) (menuList []models.SysMenu, err error) {
	var allMenus []models.SysMenu
	treeMap := make(map[int][]models.SysMenu)
	err = global.DB.Model(&models.SysMenu{}).Where(where).Order("sort desc").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}

	menuList = treeMap[0]

	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, err
}

//获取基础菜单加子菜单
func getBaseChildrenList(menu *models.SysMenu, treeMap map[int][]models.SysMenu) (err error) {

	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//获取全部菜单
func GetMenu() tool.M {
	var allMenus []models.SysMenu
	if err := global.DB.Order("sort desc").Find(&allMenus).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return tool.DataReturn(false, "查询失败", err)
	} else {
		return tool.DataReturn(true, "查询成功", allMenus)
	}
}

//获取角色的权限菜单列表
func GetRuleAuthorityMenuList(rule_id string) tool.M {
	var (
		allMenus []models.SysMenu
		err      error
	)
	where := tool.M{"hidden": 0}
	if allMenus, err = getBaseMenuTreeMap(where); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return tool.DataReturn(false, "查询失败", err)
	} else {
		formatRuleMenus(allMenus, rule_id)
		return tool.DataReturn(true, "查询成功", allMenus)
	}
}

//格式化角色菜单树
func formatRuleMenus(allMenus []models.SysMenu, rule_id string) []models.SysMenu {
	for k, v := range allMenus {
		allMenus[k].Checked = HasPolicyByRuleIdPath(rule_id, v.Path, "get|post")
		if len(v.Children) > 0 {
			allMenus[k].Children = formatRuleMenus(v.Children, rule_id)
		}
	}

	return allMenus
}

//格式化左侧菜单
func formatAllMenus(allMenus []models.SysMenu, rule_id ...string) (newallMenus []models.SysMenu) {
	for _, v := range allMenus {
		v.Checked = HasPolicyByRuleIdsPath(rule_id, v.Path, "get|post")
		if v.Checked {
			if len(v.Children) > 0 {
				v.Children = formatAllMenus(v.Children, rule_id...)
			}
			newallMenus = append(newallMenus, v)
		}
	}

	return newallMenus
}

//查询角色信息
func GetMenuInfo(id string, field ...string) tool.M {
	var menu models.SysMenu

	if err := global.DB.Select(field).First(&menu, id).Error; err != nil {
		return tool.DataReturn(false, "查无数据", err)
	}
	return tool.DataReturn(true, "查询成功", menu)
}

//获取基础菜单列表
func GetLeftMenuList(user_id string) tool.M {
	var (
		allMenus []models.SysMenu
		err      error
	)

	rule_id := GetRulesForUser(user_id)
	where := tool.M{"hidden": 0, "is_view": 1}
	if allMenus, err = getBaseMenuTreeMap(where); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return tool.DataReturn(false, "查询失败", err)
	} else {
		allMenus = formatAllMenus(allMenus, rule_id...)
		return tool.DataReturn(true, "查询成功", allMenus)
	}
}

//创建menu
func CreateChildMenu(parent_id int, menu models.SysMenu, field ...string) tool.M {
	var menuBox models.SysMenu
	//检测parent_id是否可用
	if err := global.DB.Select("id").Where("id", parent_id).First(&menuBox).Error; err != nil {
		return tool.DataReturn(false, "父菜单不存在", err)
	}
	//创建menu
	return CreateBaseMenu(menu, field...)
}

//创建基础Menu
func CreateBaseMenu(menu models.SysMenu, field ...string) tool.M {
	var (
		menuBox models.SysMenu
		err     error
	)
	//先通过name查询菜单是否存在
	err = global.DB.Select("id").Where("name = ?", menu.Name).First(&menuBox).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tool.DataReturn(false, "菜单已存在", err)
	}
	if err := global.DB.Omit(field...).Create(&menu).Error; err != nil {
		return tool.DataReturn(false, "创建失败", err)
	}
	return tool.DataReturn(true, "创建成功", nil)
}

//更新角色
func UpdateMenu(id int, menu models.SysMenu, field ...string) tool.M {
	if err := global.DB.Select("id").Where("id", id).First(&models.SysMenu{}).Error; err != nil {
		return tool.DataReturn(false, "查无菜单信息", nil)
	}

	result := global.DB.Omit(field...).Model(&menu).Updates(menu)
	if result.Error == nil && result.RowsAffected > 0 {
		return tool.DataReturn(true, "更新成功", nil)
	}
	return tool.DataReturn(false, "更新失败", result.Error)
}

//删除角色
func DelMenu(menu models.SysMenu) tool.M {
	if err := global.DB.Where("id", menu.ID).First(&models.SysMenu{}).Delete(&menu).Error; err != nil {
		return tool.DataReturn(false, "删除失败", err)
	}

	return tool.DataReturn(true, "删除成功", nil)
}
