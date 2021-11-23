package tpl

func ControllerTemplate() string {
	return `
package {{.Package}}

import (
	"go-api/{{.ModelImport}}"
	{{if ne .ServiceImport ""}}
	"go-api/{{.ServiceImport}}"
	{{end}}
	"go-api/app/services"
	"go-api/app/request"
	"go-api/tool"
	
	"github.com/gin-gonic/gin"
)

// @Summary {{.Name}}列表视图
// @Description  {{.Name}}列表视图
// @Router /admin/{{.ViewName}}View [get] 
func {{.ControllerName}}Views(c *gin.Context) {
	uid := c.GetString("uid")
	view_route := c.Request.URL.RequestURI()
	data := tool.M{
		"dataUrl":    "/admin/get{{.ControllerName}}s",
		"dataMethod": "POST",
		"addUrl":     services.GetButtonPermission(uid, view_route, "/admin/create{{.ControllerName}}"),
		"editUrl":    services.GetButtonPermission(uid, view_route, "/admin/update{{.ControllerName}}/:id"),
		"delUrl":     services.GetButtonPermission(uid, view_route, "/admin/delete{{.ControllerName}}/:id"),
	}
	tool.HTML(c, "{{.ViewName}}/{{.ViewName}}_list.html", data)
}

// @Summary {{.Name}}列表
// @Description  {{.Name}}列表
// @Router /admin/get{{.ControllerName}}s [post]
func Get{{.ControllerName}}s(c *gin.Context) {
	maps := make(map[string]interface{})
	ret := {{.ServicePackage}}.Get{{.ServiceName}}List(maps, tool.DefaultGetOffset(c), 10)
	tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
}

// @Summary 新建{{.Name}}
// @Description  新建{{.Name}}
// @Router /admin/create{{.ControllerName}} [post]
func Create{{.ControllerName}}(c *gin.Context) {

	var {{.ControllerName}} {{.ModelStruct}}
	c.ShouldBind(&{{.ControllerName}})
	//AddVerify 自行替换
	if err := request.Verify({{.ControllerName}}, request.AddVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := {{.ServicePackage}}.Create{{.ServiceName}}({{.ControllerName}}); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}

// @Summary 更新{{.Name}}信息
// @Description  更新{{.Name}}信息
// @Router /admin/update{{.ControllerName}}/:id [post]
func Update{{.ControllerName}}(c *gin.Context) {
	var {{.ControllerName}} {{.ModelStruct}}

	c.ShouldBind(&{{.ControllerName}})

	id := c.Param("id")
	{{.ControllerName}}.ID = tool.StringToInt(id)
	//UpVerify 自行替换
	if err := request.Verify({{.ControllerName}}, request.UpVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := {{.ServicePackage}}.Update{{.ServiceName}}ById(id, {{.ControllerName}}); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}

}

// @Summary 删除{{.Name}}
// @Description  删除{{.Name}}
// @Router /admin/delete{{.ControllerName}}/:id [post]
func Delete{{.ControllerName}}(c *gin.Context) {
	var {{.ControllerName}} {{.ModelStruct}}
	c.ShouldBindUri(&{{.ControllerName}})
	//DelVerify 自行替换
	if err := request.Verify({{.ControllerName}}, request.DelVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := {{.ServicePackage}}.Del{{.ServiceName}}ById({{.ControllerName}}); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}
	

`
}
