package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go-api/config"
)

type mysqlConfig struct {
	dbType        string
	dbName        string
	dbUser        string
	dbPassword    string
	dbHost        string
	dbTablePrefix string
	LogMode       bool
	MaxLifetime   time.Duration
}

var (
	db       *gorm.DB
	dbConfig = &mysqlConfig{
		LogMode:       config.MysqlSetting.LogMode,
		dbType:        "mysql",
		dbName:        config.MysqlSetting.MysqlName,
		dbUser:        config.MysqlSetting.MysqlUser,
		dbPassword:    config.MysqlSetting.MysqlPassword,
		dbHost:        config.MysqlSetting.MysqlHost,
		dbTablePrefix: config.MysqlSetting.MysqlPrefix,
		MaxLifetime:   config.MysqlSetting.MaxLifetime,
	}
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedAt  int `json:"deleted_at" gorm:"-"`
}

func init() {
	var (
		err error
	)
	db, err = gorm.Open(dbConfig.dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.dbUser,
		dbConfig.dbPassword,
		dbConfig.dbHost,
		dbConfig.dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbConfig.dbTablePrefix + defaultTableName
	}

	/* f, _ := os.Create("mysql.log")
	db.SetLogger(log.New(f, "\r\n", 0))
	mysql日志插入日志文件
	*/
	db.SingularTable(true)                           //全局禁用表名复数
	db.LogMode(dbConfig.LogMode)                     //true  打印Log
	db.DB().SetMaxIdleConns(10)                      //最大空闲连接
	db.DB().SetMaxOpenConns(100)                     //最大数据库链接
	db.DB().SetConnMaxLifetime(dbConfig.MaxLifetime) //数据库链接最大生存时间
	//
}

func State() string {
	return fmt.Sprintf("%+v", db.DB().Stats())
}

func CloseDB() {
	defer db.Close()
}

func Exec(sql string) error {
	if err := db.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
