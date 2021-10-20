package request

var (
	RuleAddVerify = Rules{"Role_name": "required", "Role_desc": "required"}
	RuleUpVerify  = Rules{"ID": "required,numeric,min=1", "Role_name": "omitempty,required", "Role_desc": "omitempty,required", "Status": "omitempty,oneof=0 1"}
	RuleDelVerify = Rules{"ID": "required,numeric,min=1"}
)
