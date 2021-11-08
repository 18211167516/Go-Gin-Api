package models

type SysMenu struct {
	Model
	ParentId int       `desc:"父菜单ID" json:"parent_id" form:"parent_id" gorm:"default:0;comment:父菜单ID"`
	Path     string    `desc:"菜单path" json:"path" form:"path" gorm:"comment:路由name"`
	Name     string    `desc:"菜单name" json:"name" form:"name" gorm:"comment:父菜单ID;"`
	Hidden   *string   `desc:"是否隐藏" json:"hidden" form:"hidden" gorm:"comment:是否在列表隐藏"`
	Sort     int       `desc:"菜单排序值" json:"sort" form:"sort" gorm:"default:1;comment:排序标记"`
	Is_view  string    `desc:"是否为视图" json:"is_view" form:"is_view" gorm:"default:0;comment:是否视图"`
	Children []SysMenu `json:"children" gorm:"-"`
	Checked  bool      `json:"checked" gorm:"-"`
}
