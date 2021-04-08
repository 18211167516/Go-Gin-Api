package models

import (
	"time"

	"gorm.io/gorm"
)

type AdminRule struct {
	CreatedOn    int64          `json:"created_on"`
	ModifiedOn   int64          `json:"modified_on"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Role_name    string         `desc:"角色名称" json:"role_name" form:"role_name" `
	Role_desc    string         `desc:"角色描述" json:"role_desc" form:"role_desc" `
	Status       int            `desc:"角色状态" json:"status"`
	Authority_id int            `desc:"角色ID" gorm:"primary_key" uri:"id" form:"id" json:"id"`
}

func (model *AdminRule) BeforeCreate(tx *gorm.DB) error {
	model.CreatedOn = time.Now().Unix()

	return nil
}

func (model *AdminRule) BeforeUpdate(tx *gorm.DB) error {
	model.ModifiedOn = time.Now().Unix()
	return nil
}
