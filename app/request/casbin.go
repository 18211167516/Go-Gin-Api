package request

type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// Casbin structure for input parameters
type Casbins struct {
	RuleId      string       `json:"ruldId"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}
