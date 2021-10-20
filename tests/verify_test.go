package tests

import (
	"fmt"
	"go-api/app/models"
	"go-api/app/request"
	"testing"
)

func TestVerify1(t *testing.T) {

	st := models.SysUser{}
	st.ID = 1
	st.Name = "我是描述"
	st.Password = "B12"

	//err := request.ValidateVariable()
	err := request.Verify(st, request.UserChangePasswordVerify)
	fmt.Println("错误", err)
	/* typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // 获取reflect.Type类型

	kd := val.Kind() // 获取到st对应的类别
	fmt.Println(typ, val, kd)

	fmt.Println(val.FieldByName("Rule_aa")) */
}
