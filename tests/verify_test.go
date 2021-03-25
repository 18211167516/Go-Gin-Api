package tests

import (
	"fmt"
	"go-api/app/request"
	"testing"
)

func TestVerify1(t *testing.T) {
	st := struct {
		Role_name string `desc:"角色名称"`
		Role_desc string `desc:"角色描述"`
	}{
		Role_name: "我是教师",
		Role_desc: "",
	}

	//err := request.ValidateVariable()
	err := request.Verify(st, request.RuleAddVerify)
	fmt.Println("错误", err)
	/* typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // 获取reflect.Type类型

	kd := val.Kind() // 获取到st对应的类别
	fmt.Println(typ, val, kd)

	fmt.Println(val.FieldByName("Rule_aa")) */
}
