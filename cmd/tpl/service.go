package tpl

func ServiceTemplate() string {
	return `
package {{.Package}}

import (
	"go-api/{{.ModelImport}}"
	"go-api/global"
	"go-api/tool"
)

//获取列表
func Get{{.ModelName}}List(maps interface{}, offset int, limit int) tool.M {
	data := make(map[string]interface{})
	var (
		{{.ModelName}}  []{{.ModelStruct}}
		count int64
	)
	global.DB.Where(maps).Limit(limit).Offset(offset).Find(&{{.ModelName}})
	global.DB.Model(&{{.ModelStruct}}{}).Where(maps).Count(&count)
	data["lists"] = {{.ModelName}}
	data["total"] = count
	if len({{.ModelName}}) > 0 {
		return tool.DataReturn(true, "查询成功", data)
	}
	return tool.DataReturn(false, "暂无数据", nil)
}

//创建
func Create{{.ModelName}}({{.ModelName}} {{.ModelStruct}}) tool.M {
	if err := global.DB.Create(&{{.ModelName}}).Error;err!=nil{
		return tool.DataReturn(false, "创建失败", err)
	}else{
		return tool.DataReturn(true, "创建成功", nil)
	}
}

//查询
func Get{{.ModelName}}ById(id string, field ...string) tool.M {
	var {{.ModelName}} {{.ModelStruct}}

	if err := global.DB.Select(field).First(&{{.ModelName}}, id).Error; err != nil {
		return tool.DataReturn(false, "查无数据", err)
	}
	return tool.DataReturn(true, "查询成功", {{.ModelName}})
}

//更新
func Update{{.ModelName}}ById(id string, {{.ModelName}} {{.ModelStruct}}) tool.M {
	result := global.DB.Where("id", id).First(&{{.ModelStruct}}{}).Updates(&{{.ModelName}})
	if result.Error == nil && result.RowsAffected > 0 {
		return tool.DataReturn(true, "更新成功", nil)
	}
	return tool.DataReturn(false, "更新失败", result.Error)
}

//删除
func Del{{.ModelName}}ById({{.ModelName}} {{.ModelStruct}}) tool.M {
	if err := global.DB.Where("id", {{.ModelName}}.ID).First(&{{.ModelStruct}}{}).Delete(&{{.ModelName}}).Error; err != nil {
		return tool.DataReturn(false, "删除失败", err)
	}
	return tool.DataReturn(true, "删除成功", nil)
}

`
}
