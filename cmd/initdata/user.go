package initdata

import (
	"go-api/app/models"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

func InitSysUsers(db *gorm.DB) {
	var status = new(int)
	*status = 1

	var Users = []models.SysUser{
		{Model: models.Model{ID: 1, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "admin", Password: "e64b78fc3bc91bcbc7dc232ba8ec59e0", RealName: "超级管理员", Status: status, Type: 1},
		{Model: models.Model{ID: 2, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Name: "test", Password: "e64b78fc3bc91bcbc7dc232ba8ec59e0", RealName: "测试同学", Status: status, Type: 2},
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]models.SysUser{}).RowsAffected == 2 {
			log.Println("sys_users表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&Users).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		log.Printf("[Mysql]--> sys_users 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
