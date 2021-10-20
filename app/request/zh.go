package request

import (
	"errors"
	"go-api/global"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

type Rules map[string]string

func init() {
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, _ = uni.GetTranslator("zh")
	va, ok := binding.Validator.Engine().(*validator.Validate)
	global.Verify = va
	if ok {
		// 验证器注册翻译器
		zh_translations.RegisterDefaultTranslations(global.Verify, trans)
		// 自定义错误信息
		global.Verify.RegisterTranslation("lowercase", trans, func(ut ut.Translator) error {
			return ut.Add("lowercase", "{0}必须为小写字母", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("lowercase", fe.Field())
			return t
		})
		// 自定义验证方法
		global.Verify.RegisterValidation("checkpassword", checkPassword)

		global.Verify.RegisterTranslation("checkpassword", trans, func(ut ut.Translator) error {
			return ut.Add("checkpassword", "规则错误", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("checkpassword", fe.Field())
			return t
		})
		global.Verify.RegisterValidation("checkMobile", checkMobile)

		global.Verify.RegisterTranslation("checkMobile", trans, func(ut ut.Translator) error {
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
	switch errs := err.(type) {
	case validator.ValidationErrors:
		return Translate(errs)
	default:
		return errs.Error()
	}

}

func VerifyMap(st map[string]interface{}, r Rules) error {
	va := global.Verify
	for k, v := range st {
		// 查看规则是否有该字段无则忽略
		verify, ok := r[k]

		if !ok {
			continue
		}

		if err := va.Var(v, verify); err != nil {
			var b strings.Builder
			b.WriteString(k)
			b.WriteString(GetError(err))
			return errors.New(b.String())
		}
	}
	return nil
}

func Verify(st interface{}, r Rules) (err error) {
	va := global.Verify
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)
	kd := val.Kind()
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	for i := 0; i < num; i++ {
		// 取每个字段
		typVal := typ.Field(i)
		valVal := val.Field(i)

		if val.Field(i).Kind() == reflect.Struct {
			err = Verify(valVal.Interface(), r)
			if err != nil {
				return err
			}
		}
		// 获取tag
		tag := typVal.Tag.Get("desc")
		// 查看规则是否有该字段无则忽略
		verify, ok := r[typVal.Name]
		if !ok {
			continue
		}
		// 获取字段的值信息
		if err = va.Var(valVal.Interface(), verify); err != nil {
			var b strings.Builder
			if tag != "" {
				b.WriteString(tag)
			} else {
				b.WriteString(typVal.Name)
			}
			b.WriteString(GetError(err))
			return errors.New(b.String())
		}
	}
	return err
}

//用于俩变量判断值
func VerifyValue(field, other interface{}, tag string) error {
	return global.Verify.VarWithValue(field, other, tag)
}

func checkMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	return len(mobile) != 11
}

func checkPassword(f1 validator.FieldLevel) bool {
	password := f1.Field().String()
	reg, _ := regexp.MatchString(`^[A-Z]+[a-z0-9]{7,17}$`, password)
	return reg
}
