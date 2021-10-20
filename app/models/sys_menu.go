package models

type SysMenu struct {
	Model
	ParentId int       `desc:"父菜单ID" json:"parent_id" form:"parent_id" gorm:"default:0"`
	Path     string    `desc:"菜单path" json:"path" form:"path" `
	Name     string    `desc:"菜单name" json:"name" form:"name" `
	Hidden   *string   `desc:"是否隐藏" json:"hidden" form:"hidden" gorm:"default:0"`
	Sort     *int      `desc:"菜单排序值" json:"sort" form:"sort"`
	Is_view  string    `desc:"是否为视图" json:"is_view" form:"is_view"`
	Children []SysMenu `json:"children" gorm:"-"`
	Checked  bool      `json:"checked" gorm:"-"`
}
