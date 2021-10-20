package request

var (
	LoginVerify = Rules{"Name": "required", "Password": "required,checkpassword", "Captcha": "required", "CaptchaId": "required"}
)

type Login struct {
	Name      string `desc:"用户名" form:"name"`
	Password  string `desc:"密码" form:"password"`
	Captcha   string `desc:"验证码" form:"captcha"`
	CaptchaId string `desc:"验证码ID" form:"captcha_id"`
}
