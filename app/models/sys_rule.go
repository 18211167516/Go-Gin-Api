package models

import (
	"fmt"
	"math/rand"
)

type SysRule struct {
	Model
	Role_name string `desc:"角色名称" json:"role_name" form:"role_name" gorm:"comment:角色名称;not null"`
	Role_desc string `desc:"角色描述" json:"role_desc" form:"role_desc" gorm:"comment:角色描述;not null"`
	Status    *int   `desc:"角色状态" json:"status" form:"status" gorm:"type:tinyint;size:1;default:1;comment:角色状态;"`
}

func (SysRule) DynamicTableName() string {
	return fmt.Sprintf("sys_rules_%d", rand.Intn(100))
}
