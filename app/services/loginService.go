package services

import (
	"errors"
	"go-api/app/models"
	"go-api/app/request"
	"go-api/app/response"
	"go-api/global"
	"go-api/tool"

	"gorm.io/gorm"
)

func Login(user request.Login) tool.M {
	var userBox response.SysLoginUserResponse
	//先通过name查询用户是否存在
	if err := global.DB.Model(&models.SysUser{}).Where("name = ? ", user.Name).First(&userBox).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tool.DataReturn(false, "查无账号", err)
		}
		return tool.DataReturn(false, "数据库错误", err)
	}

	if tool.MD5([]byte(user.Password)) != userBox.Password {
		return tool.DataReturn(false, "用户或密码错误", nil)
	}
	return tool.DataReturn(true, "登录成功", userBox)
}
