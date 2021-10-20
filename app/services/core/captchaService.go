package core

import (
	"go-api/global"
	"go-api/tool"

	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

//生成图片验证码
func Captcha() tool.M {
	driver := base64Captcha.NewDriverDigit(
		global.VP.GetInt("captcha.imgHeight"),
		global.VP.GetInt("captcha.ImgWidth"),
		global.VP.GetInt("captcha.KeyLen"),
		global.VP.GetFloat64("captcha.MaxSkew"),
		global.VP.GetInt("captcha.DotCount"))
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		return tool.DataReturn(false, "生成失败", err)
	} else {
		return tool.DataReturn(true, "生成成功", map[string]string{"id": id, "base64": b64s})
	}
}

//验证图片验证码
func Verify(id string, answer string) bool {
	return store.Verify(id, answer, true)
}
