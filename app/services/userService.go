package services

import (
	"go-api/app/models"
	"go-api/tool"
)

type UserContract interface {
	UserGetsContract    //GET users
	UserGetContract     //GET user
	UserGetByIDContract //GETByID user
	UserAddContract     //Add user
	UserExistContract   //exist user
	UserDeleteContract  //delete user
	UserEditContract    //edit user
}

type UserGetsContract interface {
	GetUserList(maps interface{}, offset int, limit int) tool.M
}

type UserGetContract interface {
	GetUser(maps interface{}) tool.M
}

type UserGetByIDContract interface {
	GetUserByID(id int) tool.M
}

type UserAddContract interface {
	AddUser(maps map[string]interface{}, data map[string]interface{}) tool.M
}

type UserExistContract interface {
	ExistUser(maps map[string]interface{}) bool
}

type UserDeleteContract interface {
	DeleteUser(maps map[string]interface{}) tool.M
}

type UserEditContract interface {
	EditUser(id int, data interface{}) tool.M
}

type UserService struct {
}

func (T UserService) GetUserList(maps interface{}, offet int, limit int) tool.M {
	data := make(map[string]interface{})
	data["lists"] = models.GetUsers(offet, limit, maps)
	data["total"] = models.GetUserTotal(maps)
	return tool.DataReturn(true, "查询成功", data)
}

func (T UserService) GetUser(maps interface{}) tool.M {
	data, err := models.GetUser(maps)
	if err != nil {
		return tool.DataReturn(false, "暂无数据", err.Error())
	}
	return tool.DataReturn(true, "查询成功", data)
}

func (T UserService) GetUserByID(id int) tool.M {
	data, err := models.GetUserByID(id)
	if err != nil {
		return tool.DataReturn(false, "暂无数据", err.Error())
	}
	return tool.DataReturn(true, "查询成功", data)
}

func (T UserService) AddUser(maps map[string]interface{}, data map[string]interface{}) tool.M {
	//先通过maps查询数据是否存在
	if T.ExistUser(maps) {
		return tool.DataReturn(false, "该名称已存在", nil)
	} else {
		if isbool := models.AddUser(data); !isbool {
			return tool.DataReturn(false, "创建失败", nil)
		}
		return tool.DataReturn(true, "创建成功", nil)
	}
}

func (T UserService) ExistUser(maps map[string]interface{}) bool {

	return models.ExistUserByMaps(maps)
}

func (T UserService) DeleteUser(maps map[string]interface{}) tool.M {
	if T.ExistUser(maps) {
		isbool, err := models.DeleteUser(maps)
		if err != nil {
			return tool.DataReturn(false, "删除失败", err.Error())
		}
		return tool.DataReturn(isbool, "删除成功", nil)
	} else {
		return tool.DataReturn(false, "记录不存在", nil)
	}
}

func (T UserService) EditUser(id int, data interface{}) tool.M {
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
