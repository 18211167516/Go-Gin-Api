package services

import (
	"go-api/app/models"
	"go-api/global"
	"go-api/tool"
)

func GetRoleList(maps interface{}, offset int, limit int) tool.M {
	data := make(map[string]interface{})
	var (
		rule  []models.AdminRule
		count int64
	)
	global.DB.Where(maps).Limit(limit).Offset(offset).Find(&rule)
	global.DB.Model(&models.AdminRule{}).Where(maps).Count(&count)
	data["lists"] = rule
	data["total"] = count
	return tool.DataReturn(true, "查询成功", data)
}

func CreateRule(rule models.AdminRule) tool.M {
	//先通过maps查询数据是否存在
	if err := global.DB.Create(&rule).Error; err != nil {
		return tool.DataReturn(false, "创建失败", err)
	}
	return tool.DataReturn(true, "创建成功", nil)
}

func UpdateRule(rule models.AdminRule) tool.M {
	result := global.DB.Where("authority_id", rule.Authority_id).First(&models.AdminRule{}).Updates(&rule)
	if result.Error == nil && result.RowsAffected > 0 {
		return tool.DataReturn(true, "更新成功", nil)
	}
	return tool.DataReturn(false, "更新失败", result.Error)

}

func DelRule(rule models.AdminRule) tool.M {
	if err := global.DB.Where("authority_id", rule.Authority_id).First(&models.AdminRule{}).Delete(&rule).Error; err != nil {
		return tool.DataReturn(false, "删除失败", err)
	}

	return tool.DataReturn(true, "删除成功", nil)
}
