package models

import (
    "log"
    "fmt"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    
    "go-api/config"
)

type mysqlConfig struct{
	dbType string
	dbName string
	dbUser string
	dbPassword string
	dbHost string
	dbTablePrefix string
	LogMode bool
} 

var (
	db *gorm.DB
	dbConfig = &mysqlConfig{
		dbType:"mysql",
		dbName:config.MysqlSetting.MysqlName,
		dbUser:config.MysqlSetting.MysqlUser,
		dbPassword:config.MysqlSetting.MysqlPassword,
		dbHost:config.MysqlSetting.MysqlHost,
		dbTablePrefix:config.MysqlSetting.MysqlPrefix,
	}
)

type Model struct {
    ID int `gorm:"primary_key" json:"id"`
    CreatedOn int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
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

    gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
        return dbConfig.dbTablePrefix + defaultTableName;
    }

    db.SingularTable(true)
    db.LogMode(dbConfig.LogMode)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
    defer db.Close()
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())

    return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())
    return nil
}