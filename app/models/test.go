package models

type Test struct {
	Model
	Name string `desc:"菜单name" json:"name" form:"name" gorm:"comment:名称;not null"`
	Sort int    `desc:"菜单排序值" json:"sort" form:"sort" gorm:"type:int;size:10;default:1;comment:排序标记;not null"`
}

func init() {
	AutoMigratFunc["test"] = func() interface{} {
		return Test{}
	}
}
