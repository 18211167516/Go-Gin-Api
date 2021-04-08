package models

type SysMenu struct {
	Model
	MenuLevel uint      `json:"-"`
	ParentId  string    `json:"parentId" form:"parentId" gorm:"comment:父菜单ID"`
	Path      string    `json:"path" form:"path" gorm:"comment:路由path"`
	Name      string    `json:"name" form:"name" gorm:"comment:路由name"`
	Hidden    bool      `json:"hidden" form:"hidden" gorm:"comment:是否在列表隐藏"`
	Component string    `json:"component" form:"component" gorm:"comment:对应前端文件路径"`
	Sort      int       `json:"sort" form:"sort" gorm:"comment:排序标记"`
	Children  []SysMenu `json:"children" gorm:"-"`
}
