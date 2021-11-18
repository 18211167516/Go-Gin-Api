package models

type SysMenu struct {
	Model
	ParentId int       `desc:"父菜单ID" json:"parent_id" form:"parent_id" gorm:"index;default:0;comment:父菜单ID"`
	Path     string    `desc:"菜单path" json:"path" form:"path" gorm:"comment:路由地址;not null"`
	Name     string    `desc:"菜单name" json:"name" form:"name" gorm:"comment:路由名称;not null"`
	Hidden   *string   `desc:"是否隐藏" json:"hidden" form:"hidden" gorm:"type:tinyint;size:1;comment:是否在列表隐藏;not null"`
	Sort     int       `desc:"菜单排序值" json:"sort" form:"sort" gorm:"type:int;size:10;default:1;comment:排序标记;not null"`
	Is_view  string    `desc:"是否为视图" json:"is_view" form:"is_view" gorm:"type:tinyint;size:1;default:0;comment:是否视图;not null"`
	Children []SysMenu `json:"children" gorm:"-"`
	Checked  bool      `json:"checked" gorm:"-"`
}
