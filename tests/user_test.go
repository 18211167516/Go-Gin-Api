package tests

import (
	"go-api/app/models"
	"testing"
)

func testGetUsers(t *testing.T) {

	testcase := []TestCase{
		{
			code:         0,
			param:        `name=5678999999`,
			errMsg:       `查询成功`,
			method:       "GET",
			desc:         "验证查询用户列表成功",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "",
			ext1:         1,
		},
		{
			code:         40001,
			param:        ``,
			errMsg:       `ID为必填字段`,
			method:       "GET",
			desc:         "验证ID不能为空",
			showBody:     true,
			url:          "/api/v1/user/0",
			content_type: "",
			ext1:         1,
		},
		{
			code:         0,
			param:        `name=5678999999`,
			errMsg:       `查询成功`,
			method:       "GET",
			desc:         "验证查询用户ID1成功",
			showBody:     true,
			url:          "/api/v1/user/1",
			content_type: "",
			ext1:         1,
		},
		{
			code:         40001,
			param:        ``,
			errMsg:       `暂无数据`,
			method:       "GET",
			desc:         "验证查询用户ID10000不存在",
			showBody:     true,
			url:          "/api/v1/user/1000",
			content_type: "",
			ext1:         1,
		},
	}

	call(t, testcase)
}

func testAddUsers(t *testing.T) {

	testcase := []TestCase{
		{
			code:         400,
			param:        `name=&created_by=admin`,
			errMsg:       `参数Name为必填字段错误`,
			method:       "POST",
			desc:         "验证name必填",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         400,
			param:        `name=123123&created_by=Admin`,
			errMsg:       `参数CreatedBy必须为小写字母错误`,
			method:       "POST",
			desc:         "验证created_by参数异常",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         400,
			param:        `{"name":"","created_by":"admin"}`,
			errMsg:       `参数Name为必填字段错误`,
			method:       "POST",
			desc:         "验证JSON格式下name参数为空",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/json",
			ext1:         1,
		},
		{
			code:         400,
			param:        `{"name":"1233","created_by":"Admin"}`,
			errMsg:       `参数CreatedBy必须为小写字母错误`,
			method:       "POST",
			desc:         "验证JSON格式下created_by参数异常",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/json",
			ext1:         1,
		},
		{
			code: 400,
			param: `<?xml version="1.0" encoding="UTF-8"?>
				<root>
					<name></name>
					<created_by>123</created_by>
				</root>`,
			errMsg:       `参数Name为必填字段错误`,
			method:       "POST",
			desc:         "验证XML格式下name参数为空",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/xml",
			ext1:         1,
		},
		{
			code: 400,
			param: `<?xml version="1.0" encoding="UTF-8"?>
			<root>
				<name>123213</name>
				<created_by>Admin</created_by>
			</root>`,
			errMsg:       `参数CreatedBy必须为小写字母错误`,
			method:       "POST",
			desc:         "验证XML格式下created_by参数异常",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/xml",
			ext1:         1,
		},
		{
			code:         0,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `创建成功`,
			method:       "POST",
			desc:         "验证创建成功",
			showBody:     true,
			url:          "/api/v1/users",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
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
	}
	call(t, testcase)

}

func testEditUsers(t *testing.T) {
	testcase := []TestCase{
		{
			code:         400,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `参数ID为必填字段错误`,
			method:       "PUT",
			desc:         "验证ID不存在",
			showBody:     true,
			url:          "/api/v1/users/0",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         400,
			param:        `name=&created_by=admin`,
			errMsg:       `参数Name为必填字段错误`,
			method:       "PUT",
			desc:         "验证name参数异常",
			showBody:     true,
			url:          "/api/v1/users/1",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         400,
			param:        `name=mam&created_by=Admin`,
			errMsg:       `参数CreatedBy必须为小写字母错误`,
			method:       "PUT",
			desc:         "验证created_by参数异常",
			showBody:     true,
			url:          "/api/v1/users/1",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         40001,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `ID记录不存在`,
			method:       "PUT",
			desc:         "验证ID记录不存在",
			showBody:     true,
			url:          "/api/v1/users/2",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
		{
			code:         0,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `编辑成功`,
			method:       "PUT",
			desc:         "验证编辑成功",
			showBody:     true,
			url:          "/api/v1/users/1",
			content_type: "application/x-www-form-urlencoded",
			ext1:         1,
		},
	}

	call(t, testcase)
}

func testDeleteUsers(t *testing.T) {
	testcase := []TestCase{
		{
			code:         400,
			param:        `name=baibai&created_by=admin`,
			errMsg:       `参数ID为必填字段错误`,
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
			errMsg:       `记录不存在`,
			method:       "DELETE",
			desc:         "验证记录不存在",
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
	}

	call(t, testcase)
}

func testTruncate(t *testing.T) {
	sql := "truncate test_user;"
	if err := models.Exec(sql); err != nil {
		t.Errorf("%s", err)
	}
}
func TestUserAll(t *testing.T) {
	t.Run("TestTruncate", testTruncate)
	t.Run("TestAddUsers", testAddUsers)
	t.Run("TestGetUsers", testGetUsers)
	t.Run("TestEditUsers", testEditUsers)
	t.Run("TestDeleteUsers", testDeleteUsers)
}
