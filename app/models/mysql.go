package models

import (
	"fmt"

	"go-api/global"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedAt  int `json:"deleted_at" gorm:"-"`
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

/* func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
} */
