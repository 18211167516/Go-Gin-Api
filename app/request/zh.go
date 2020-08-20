package request

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var v *validator.Validate
var trans ut.Translator

func init() {
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, _ = uni.GetTranslator("zh")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 验证器注册翻译器
		zh_translations.RegisterDefaultTranslations(v, trans)
		// 自定义错误信息
		v.RegisterTranslation("lowercase", trans, func(ut ut.Translator) error {
			return ut.Add("lowercase", "{0}必须为小写字母", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("lowercase", fe.Field())
			return t
		})
		// 自定义验证方法
		v.RegisterValidation("checkMobile", checkMobile)

		v.RegisterTranslation("checkMobile", trans, func(ut ut.Translator) error {
			return ut.Add("checkMobile", "{0}长度不等于11位或{1}格式错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("checkMobile", fe.Field(), fe.Field())
			return t
		})
	}
}

func Translate(errs validator.ValidationErrors) string {
	var errList []string
	for _, e := range errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(trans))

	}
	return errList[0] //strings.Join(errList, "|")
}

func GetError(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		return Translate(err.(validator.ValidationErrors))
	default:
		return errors.New("unknown error.").Error()
	}

}

func checkMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	if len(mobile) != 11 {
		return false
	}
	return true
}
