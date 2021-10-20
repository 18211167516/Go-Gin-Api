package request

var (
	MenuAddBaseVerify  = Rules{"Path": "required", "Name": "required", "Sort": "numeric,min=0", "Is_view": "oneof=0 1"}
	MenuAddChildVerify = Rules{"ParentId": "required,numeric,min=1", "Path": "required", "Name": "required", "Sort": "numeric,min=0", "Is_view": "oneof=0 1"}
	MenuUpVerify       = Rules{"ID": "required,numeric,min=1", "Name": "omitempty,required", "Path": "omitempty,required", "Sort": "omitempty,numeric,min=0", "Hidden": "omitempty,oneof=0 1", "Is_view": "omitempty,oneof=0 1"}
	MenuDelVerify      = Rules{"ID": "required,numeric,min=1"}
)
