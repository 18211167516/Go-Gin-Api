package request

var (
	CasbinViewVerify   = Rules{"ID": "required,numeric,min=1"}
	CasbinUpVerify     = Rules{"RuleId": "required,numeric,min=1", "CasbinInfos": "required"}
	UserUpCasbinVerify = Rules{"UserId": "required,numeric,min=1", "Rules": "required"}
)

type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// Casbin structure for input parameters
type Casbins struct {
	RuleId      string `desc:"角色ID" json:"rule_id" form:"rule_id" `
	CasbinInfos []int  `desc:"权限菜单" json:"rules" form:"rules[]"`
}

type CasbinsRules struct {
	UserId string   `desc:"用户ID" json:"user_id" form:"user_id"`
	Rules  []string `desc:"角色" json:"rules" form:"rules[]"`
}
