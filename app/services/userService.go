package services

import (
	"go-api/app/models"
	"go-api/tool"
)

func GetUserList(maps interface{}, offet int, limit int) tool.M {
	data := make(map[string]interface{})
	data["lists"] = models.GetUsers(offet, limit, maps)
	data["total"] = models.GetUserTotal(maps)
	return tool.DataReturn(true, "查询成功", data)
}

func GetUser(maps interface{}) tool.M {
	data, err := models.GetUser(maps)
	if err != nil {
		return tool.DataReturn(false, "暂无数据", err.Error())
	}
	return tool.DataReturn(true, "查询成功", data)
}

func GetUserByID(id int) tool.M {
	data, err := models.GetUserByID(id)
	if err != nil {
		return tool.DataReturn(false, "暂无数据", err.Error())
	}
	return tool.DataReturn(true, "查询成功", data)
}

func AddUser(maps map[string]interface{}, data map[string]interface{}) tool.M {
	//先通过maps查询数据是否存在
	if ExistUser(maps) {
		return tool.DataReturn(false, "该名称已存在", nil)
	} else {
		if isbool := models.AddUser(data); !isbool {
			return tool.DataReturn(false, "创建失败", nil)
		}
		return tool.DataReturn(true, "创建成功", nil)
	}
}

func ExistUser(maps map[string]interface{}) bool {

	return models.ExistUserByMaps(maps)
}

func DeleteUser(maps map[string]interface{}) tool.M {
	if ExistUser(maps) {
		isbool, err := models.DeleteUser(maps)
		if err != nil {
			return tool.DataReturn(false, "删除失败", err.Error())
		}
		return tool.DataReturn(isbool, "删除成功", nil)
	} else {
		return tool.DataReturn(false, "记录不存在", nil)
	}
}

func EditUser(id int, data interface{}) tool.M {
	if models.ExistTagByID(id) {
		_, err := models.EditUser(id, data)
		if err != nil {
			return tool.DataReturn(false, "编辑失败", err.Error())
		}
		return tool.DataReturn(true, "编辑成功", nil)
	} else {
		return tool.DataReturn(false, "ID记录不存在", nil)
	}
}
