package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-api/tool"
)

func testGetUsers(t *testing.T) {

	testcase := []TestCase{
		{
			code:         40001,
			param:        `name=5678999999`,
			errMsg:       `暂无数据`,
			method:       "GET",
			desc:         "验证查询成功",
			showBody:     false,
			url:          "/api/v1/user/0",
			content_type: "",
			ext1:         1,
		},
		{
			code:         0,
			param:        `name=5678999999`,
			errMsg:       `查询成功`,
			method:       "GET",
			desc:         "验证查询成功",
			showBody:     true,
			url:          "/api/v1/user/1",
			content_type: "",
			ext1:         1,
		},
	}

	call(t,testcase)
}

func testAddUsers(t *testing.T) {

	testcase := []TestCase{
		{
			code:         40001,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `该名称已存在`,
			method:       "POST",
			desc:         "验证插入名称已存在",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         0,
			param:        `name=5678999999&created_by=admin`,
			errMsg:       `创建成功`,
			method:       "POST",
			desc:         "验证创建成功",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
	}
	for k, v := range testcase {
		_, _, w := PerformRequest(v.method, v.url, v.content_type, v.param)
		//assert.Contains(t, w.Body.String(),fmt.Sprintf("\"error_code\":%d",v.code))
		fmt.Println()
		fmt.Printf("第%d个测试用例：%s", k+1, v.desc)
		if v.showBody {
			fmt.Printf("接口返回%s", w.Body.String())
			fmt.Println()
		}

		s := struct {
			Error_code int         `json:"error_code"`
			Msg        string      `json:"msg"`
			Data       interface{} `json:"data"`
		}{}
		err := tool.JsonToStruct([]byte(w.Body.String()), &s)
		assert.NoError(t, err)
		assert.Equal(t, v.errMsg, s.Msg, "错误信息不一致")
		assert.Equal(t, v.code, s.Error_code, "错误码不一致")
	}

}

func testEditUsers(t *testing.T) {
	testcase := []TestCase{
		{
			code:         40001,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `ID不存在`,
			method:       "PUT",
			desc:         "验证ID不存在",
			showBody:     true,
			url:          "/api/v1/users/0",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         40001,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `编辑失败`,
			method:       "PUT",
			desc:         "验证编辑失败",
			showBody:     true,
			url:          "/api/v1/users/2",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         0,
			param:        `name=mam&created_by=admin`,
			errMsg:       `编辑成功`,
			method:       "PUT",
			desc:         "验证编辑成功",
			showBody:     true,
			url:          "/api/v1/users/2",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
	}

	call(t,testcase)
}

func testDeleteUsers(t *testing.T) {
	testcase := []TestCase{
		{
			code:         40001,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `ID不存在`,
			method:       "DELETE",
			desc:         "验证ID不存在",
			showBody:     true,
			url:          "/api/v1/users/0",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         40001,
			param:        ``,
			errMsg:       `删除失败`,
			method:       "DELETE",
			desc:         "验证删除失败",
			showBody:     true,
			url:          "/api/v1/users/10000",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         0,
			param:        `name=mam&created_by=admin`,
			errMsg:       `删除成功`,
			method:       "DELETE",
			desc:         "验证删除成功",
			showBody:     true,
			url:          "/api/v1/users/1",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         40001,
			param:        `name=mam&created_by=admin`,
			errMsg:       `信息不存在`,
			method:       "DELETE",
			desc:         "验证信息不存在",
			showBody:     true,
			url:          "/api/v1/users/2",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
	}

	call(t,testcase)
}
func TestUserAll(t *testing.T) {
	t.Run("TestGetUsers", testGetUsers)
	//t.Run("TestAddUsers", testAddUsers)
	//t.Run("TestEditUsers", testEditUsers)
	//t.Run("TestDeleteUsers", testDeleteUsers)
}
