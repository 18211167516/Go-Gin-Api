package request

var (
	RuleAddVerify = Rules{"Role_name": "required", "Role_desc": "required"}
	RuleUpVerify  = Rules{"Authority_id": "required", "Role_name": "omitempty,required", "Role_desc": "omitempty,required"}
	RuleDelVerify = Rules{"Authority_id": "required"}
)
