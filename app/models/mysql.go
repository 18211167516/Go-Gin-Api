package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"go-api/global"

	"gorm.io/gorm"
)

type XTime struct {
	time.Time
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type Model struct {
	ID        int            `gorm:"primary_key" json:"id" uri:"id"`
	CreatedAt XTime          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt XTime          `json:"updated_at" gorm:"autoUpdateTime" `
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type DynamicTableName interface {
	DynamicTableName() string
}

//设置动态表名
func DynamicTable(d DynamicTableName) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(d.DynamicTableName())
	}
}

func State() string {
	sqlDB, _ := global.DB.DB()

	return fmt.Sprintf("%+v", sqlDB.Stats())
}

func Exec(sql string) error {
	if err := global.DB.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
