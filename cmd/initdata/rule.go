package initdata

import (
	"go-api/app/models"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

func InitSysRules(db *gorm.DB) {
	var status = new(int)
	*status = 1

	var rules = []models.SysRule{
		{Model: models.Model{ID: 39, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Role_name: "基础权限", Role_desc: "基础权限", Status: status},
		{Model: models.Model{ID: 40, CreatedAt: models.XTime{time.Now()}, UpdatedAt: models.XTime{time.Now()}}, Role_name: "超级管理员", Role_desc: "全部权限", Status: status},
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{39, 40}).Find(&[]models.SysRule{}).RowsAffected == 2 {
			log.Println("sys_rules表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&rules).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		log.Printf("[Mysql]--> sys_rules 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
