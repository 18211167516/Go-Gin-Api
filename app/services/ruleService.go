package services

import (
	"errors"
	"go-api/app/models"
	"go-api/app/response"
	"go-api/global"
	"go-api/tool"

	"gorm.io/gorm"
)

//获取角色列表
func GetRoleList(maps interface{}, offset int, limit int) tool.M {
	data := make(map[string]interface{})
	var (
		rule  []models.SysRule
		count int64
	)
	global.DB.Where(maps).Limit(limit).Offset(offset).Find(&rule)
	global.DB.Model(&models.SysRule{}).Where(maps).Count(&count)
	data["lists"] = rule
	data["total"] = count
	if len(rule) > 0 {
		return tool.DataReturn(true, "查询成功", data)
	}
	return tool.DataReturn(false, "暂无数据", nil)
}

//创建角色
func CreateRule(rule models.SysRule) tool.M {
	var (
		ruleBox models.SysRule
		err     error
	)
	//先通过role_name查询角色是否存在
	err = global.DB.Where("role_name = ?", rule.Role_name).First(&ruleBox).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tool.DataReturn(false, "角色已存在", err)
	}
	err = global.DB.Create(&rule).Error
	if err != nil {
		return tool.DataReturn(false, "创建失败", err)
	}
	return tool.DataReturn(true, "创建成功", nil)
}

//查询角色信息
func GetRuleInfo(id string, field ...string) tool.M {
	var rule models.SysRule

	if err := global.DB.Select(field).First(&rule, id).Error; err != nil {
		return tool.DataReturn(false, "查无数据", err)
	}
	return tool.DataReturn(true, "查询成功", rule)
}

//更新角色
func UpdateRule(id string, rule models.SysRule) tool.M {
	result := global.DB.Where("id", id).First(&models.SysRule{}).Updates(&rule)
	if result.Error == nil && result.RowsAffected > 0 {
		//如果状态是禁用的话，把该角色从casbin_rule删除
		if *rule.Status == 0 {
			ClearRule(id)
		}
		return tool.DataReturn(true, "更新成功", nil)
	}
	return tool.DataReturn(false, "更新失败", result.Error)

}

//删除角色
func DelRule(rule models.SysRule) tool.M {
	if err := global.DB.Where("id", rule.ID).First(&models.SysRule{}).Delete(&rule).Error; err != nil {
		return tool.DataReturn(false, "删除失败", err)
	}
	ClearRule(tool.IntToString(rule.ID))
	return tool.DataReturn(true, "删除成功", nil)
}

//获取用户角色列表
func GetRuleByUserid(userId string) tool.M {
	var allRules []response.SysRuleResponse
	if err := global.DB.Model(&models.SysRule{}).Where("status", 1).Find(&allRules).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return tool.DataReturn(false, "查询失败", err)
	} else {
		formatRuls(userId, allRules)
		return tool.DataReturn(true, "查询成功", allRules)
	}
}

//格式化角色列表主要获得Checked
func formatRuls(userid string, allRules []response.SysRuleResponse) {
	for k, v := range allRules {
		allRules[k].Checked = HasRuleForUser(userid, tool.IntToString(v.ID))
	}
}
