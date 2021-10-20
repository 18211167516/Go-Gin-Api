package models

type SysUser struct {
	Model
	Name            string `desc:"用户名" json:"name" form:"name" `
	RealName        string `desc:"真实姓名" json:"real_name" form:"real_name"`
	Type            int    `desc:"用户身份" json:"type" form:"type"`
	Status          *int   `desc:"用户状态" json:"status" form:"status" gorm:"default:1"`
	Password        string `desc:"用户密码" json:"Password" form:"password"  gorm:"comment:用户登录密码"`
	ConfirmPassword string `desc:"确认密码" json:"ConfirmPassword" form:"confirmPassword" gorm:"-"`
}

type UserSwagger struct {
	Lists []*SysUser
	Total int
}
