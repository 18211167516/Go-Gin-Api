package response

type SysLoginUserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"  `
	RealName string `json:"real_name"`
	Type     int    `json:"type"`
	Password string `json:"-"`
}
