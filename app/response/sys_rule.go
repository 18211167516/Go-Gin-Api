package response

import (
	"go-api/app/models"
)

type SysRuleResponse struct {
	models.SysRule
	Checked bool `json:"checked" gorm:"-"`
}
