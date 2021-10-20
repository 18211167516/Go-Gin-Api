package request

var (
	UserAddVerify               = Rules{"Path": "required", "Name": "required", "RealName": "required", "Password": "required,checkpassword", "Type": "omitempty,oneof=1 2"}
	UserUpVerify                = Rules{"ID": "required,numeric,min=1", "Name": "omitempty,required", "RealName": "omitempty,required", "Status": "omitempty,oneof=0 1", "Type": "omitempty,oneof=1 2"}
	UserChangeOwnVerify         = Rules{"ID": "required,numeric,min=1", "Name": "omitempty,required", "RealName": "omitempty,required"}
	UserChangePasswordVerify    = Rules{"ID": "required,numeric,min=1", "Password": "required,checkpassword"}
	UserChangeOwnPasswordVerify = Rules{"ID": "required,numeric,min=1", "Password": "required,checkpassword"}
	UserDelVerify               = Rules{"ID": "required,numeric,min=1"}
)

type ChangeOwnPassword struct {
	OldPassword     string `form:"old_password"`
	NewPassword     string `form:"new_password"`
	ConfirmPassword string `form:"confirm_password"`
}
