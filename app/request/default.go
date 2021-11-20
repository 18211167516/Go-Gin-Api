package request

var (
	AddVerify = Rules{}
	UpVerify  = Rules{"ID": "required,numeric,min=1"}
	DelVerify = Rules{"ID": "required,numeric,min=1"}
)
