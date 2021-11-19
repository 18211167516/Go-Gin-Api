package models

type Test struct {
	Model
	Name string `desc:"名称" gorm:"column:name;comment:名称" json:"Name"  form:"Name"`     // 名称
	Sort int64  `desc:"排序标记" gorm:"column:sort;comment:排序标记" json:"Sort"  form:"Sort"` // 排序标记
}

/*
func init() {
	AutoMigratFunc["Test"] = func() interface{} {
		return Test{}
	}
}
*/
