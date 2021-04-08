package models

import (
	"fmt"

	"go-api/global"

	"gorm.io/gorm"
)

type Model struct {
	ID         int            `gorm:"primary_key" json:"id"`
	CreatedOn  int64          `json:"created_on"`
	ModifiedOn int64          `json:"modified_on"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"-"`
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
