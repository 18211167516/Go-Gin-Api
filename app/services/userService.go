package services

import (
	"errors"
	"go-api/app/models"
	"go-api/global"
	"go-api/tool"

	"gorm.io/gorm"
)

//获取用户列表数据
func GetUserList(maps interface{}, offset int, limit int) tool.M {
	data := make(map[string]interface{})
	var (
		user  []models.SysUser
		count int64
	)

	global.DB.Where(maps).Limit(limit).Offset(offset).Find(&user)
	global.DB.Model(&models.SysUser{}).Where(maps).Count(&count)
	data["lists"] = user
	data["total"] = count
	if len(user) > 0 {
		return tool.DataReturn(true, "查询成功", data)
	}
	return tool.DataReturn(false, "暂无数据", nil)
}

//根据用户ID查询信息
func GetUserByID(id string) tool.M {
	var user models.SysUser

	if err := global.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tool.DataReturn(false, "用户不存在", err)
		}
		return tool.DataReturn(false, err.Error(), err)
	}

	return tool.DataReturn(true, "查询成功", user)
}

//添加用户
func Register(user models.SysUser) tool.M {
	var (
		userBox models.SysUser
		err     error
	)
	//先通过name查询角色是否存在
	err = global.DB.Where("name = ?", user.Name).First(&userBox).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tool.DataReturn(false, "用户名已注册", err)
	}
	user.Password = tool.MD5([]byte(user.Password))
	err = global.DB.Create(&user).Error
	if err != nil {
		return tool.DataReturn(false, "创建失败", err)
	}
	return tool.DataReturn(true, "创建成功", nil)
}

//删除用户
func DeleteUser(user models.SysUser) tool.M {
	if err := global.DB.Where("id", user.ID).First(&models.SysUser{}).Delete(&user).Error; err != nil {
		return tool.DataReturn(false, "删除失败", err)
	}

	return tool.DataReturn(true, "删除成功", nil)
}

//修改密码
func ChangePassword(id int, user models.SysUser) tool.M {
	return UpdateUser(id, user, "id", "name", "real_name", "type", "status")
}

//修改自己密码
func ChangeOwnPassword(id string, OldPassword string, newpassword string) tool.M {
	var info models.SysUser
	if err := global.DB.Select("id", "password").Where("id", id).First(&info).Error; err != nil {
		return tool.DataReturn(false, "查无用户信息", nil)
	}

	if info.Password != tool.MD5([]byte(OldPassword)) {
		return tool.DataReturn(false, "原密码错误", nil)
	}

	newpassword = tool.MD5([]byte(newpassword))

	result := global.DB.Model(&models.SysUser{}).Where("id", id).Update("password", newpassword)
	if result.Error == nil && result.RowsAffected > 0 {
		return tool.DataReturn(true, "修改密码成功", nil)
	}
	return tool.DataReturn(false, "修改密码失败", result.Error)
}

//更新用户
func UpdateUser(id int, user models.SysUser, field ...string) tool.M {
	var userBox models.SysUser
	if err := global.DB.Select("id", "name").Where("id", id).First(&userBox).Error; err != nil {
		return tool.DataReturn(false, "查无用户信息", nil)
	}

	if user.Name != "" && userBox.Name != user.Name {
		if err := global.DB.Select("id").Where("name = ?", user.Name).First(&models.SysUser{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			return tool.DataReturn(false, "用户名已注册", err)
		}
	}

	if user.Password != "" {
		user.Password = tool.MD5([]byte(user.Password))
	}
	result := global.DB.Omit(field...).Model(&user).Updates(user)
	if result.Error == nil && result.RowsAffected > 0 {
		return tool.DataReturn(true, "更新成功", nil)
	}
	return tool.DataReturn(false, "更新失败", result.Error)
}
