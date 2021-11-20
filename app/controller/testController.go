
package controller

import (
	"go-api/app/models"
	
	"go-api/app/services"
	"go-api/app/request"
	"go-api/tool"
	
	"github.com/gin-gonic/gin"
)

// @Summary 角色列表视图
// @Description  角色列表视图
// @Router /admin/TestView [get] 
func TestView(c *gin.Context) {
	uid := c.GetString("uid")
	view_route := c.Request.URL.RequestURI()
	data := tool.M{
		"dataUrl":    "/admin/getTests",
		"dataMethod": "POST",
		"addUrl":     services.GetButtonPermission(uid, view_route, "/admin/createTest"),
		"editUrl":    services.GetButtonPermission(uid, view_route, "/admin/updateTest/:id"),
		"delUrl":     services.GetButtonPermission(uid, view_route, "/admin/deleteTest/:id"),
	}
	tool.HTML(c, "test/test_list.html", data)
}

// @Summary 角色列表
// @Description  角色列表
// @Router /admin/getTests [post]
func GetTests(c *gin.Context) {
	maps := make(map[string]interface{})
	ret := services.GetTestList(maps, tool.DefaultGetOffset(c), 10)
	tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
}

// @Summary 新建角色
// @Description  新建角色
// @Router /admin/createTest [post]
func CreateTest(c *gin.Context) {

	var Test models.Test
	c.ShouldBind(&Test)
	//AddVerify 自行替换
	if err := request.Verify(Test, request.AddVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.CreateTest(Test); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}

// @Summary 更新角色信息
// @Description  更新角色信息
// @Router /admin/updateTest/:id [post]
func UpdateTest(c *gin.Context) {
	var Test models.Test

	c.ShouldBind(&Test)

	id := c.Param("id")
	Test.ID = tool.StringToInt(id)
	//UpVerify 自行替换
	if err := request.Verify(Test, request.UpVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.UpdateTestById(id, Test); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}

// @Summary 删除角色
// @Description  删除角色
// @Router /admin/deleteTest/:id [post]
func DeleteTest(c *gin.Context) {
	var Test models.Test
	c.ShouldBindUri(&Test)
	//DelVerify 自行替换
	if err := request.Verify(Test, request.DelVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.DelTestById(Test); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}
	

