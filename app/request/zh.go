package request

import (
	"errors"
	"fmt"
	"go-api/global"
	"reflect"
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
	switch err.(type) {
	case validator.ValidationErrors:
		return Translate(err.(validator.ValidationErrors))
	default:
		return err.Error()
	}

}

func Verify(st interface{}, r Rules) error {
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
		// 获取tag
		tag := typVal.Tag.Get("desc")
		// 查看规则是否有该字段无则忽略
		verify, ok := r[typVal.Name]
		if !ok {
			continue
		}
		// 获取字段的值信息
		v := val.Field(i).Interface()

		if err := va.Var(v, verify); err != nil {
			fmt.Println(GetError(err))
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
	return nil
}

func checkMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	if len(mobile) != 11 {
		return false
	}
	return true
}
