package initdata

import (
	"go-api/app/models"
	"go-api/global"
	"log"
	"os"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitMysqlData(db *gorm.DB) {
	InitSysUsers(db)
	InitSysRules(db)
	InitSysMenus(db)
	InitCasbin(db)
	log.Println("[Mysql]-->初始化数据成功")
}

func InitMysqlTables(db *gorm.DB) {
	var err error
	if !db.Migrator().HasTable("casbin_rule") {
		if _, err := gormadapter.NewAdapterByDB(global.DB); err != nil {
			log.Printf("[Mysql]-->casbin_rule初始化数据表失败,err: %v\n", err)
			os.Exit(0)
		}
	}
	err = db.AutoMigrate(
		models.SysRule{},
		models.SysUser{},
		models.SysMenu{},
	)
	if err != nil {
		log.Printf("[Mysql]-->初始化数据表失败,err: %v\n", err)
		os.Exit(0)
	}
	log.Println("[Mysql]-->初始化数据表成功")
}
