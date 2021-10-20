package controller

import (
	"go-api/app/models"
	"go-api/app/request"
	"go-api/app/services"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

// @Summary 菜单列表视图
// @Description  菜单列表视图
// @Router /admin/menusView  [get]
func MenusView(c *gin.Context) {

	list := services.GetMenu()
	data := tool.M{
		"addBaseUrl":  "/admin/createBaseMenu",
		"addChildUrl": "/admin/createChildMenu/:parent_id",
		"editUrl":     "/admin/updateMenu/:id",
		"delUrl":      "/admin/deleteMenu/:id",
		"list":        tool.StructToJson(list["data"]),
	}

	tool.HTML(c, "menu/menu_list.html", data)
}

// @Summary 创建根菜单
// @Description  创建根菜单
// @Router /admin/createBaseMenu  [post]
func CreateBaseMenu(c *gin.Context) {
	var menu models.SysMenu
	c.ShouldBind(&menu)
	if err := request.Verify(menu, request.MenuAddBaseVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.CreateBaseMenu(menu, "parent_id", "hidden"); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 创建子菜单
// @Description  创建子菜单
// @Router /admin/createChildMenu/:parent_id  [post]
func CreateChildMenu(c *gin.Context) {
	var menu models.SysMenu
	c.ShouldBind(&menu)

	parent_id := c.Param("parent_id")
	menu.ParentId = tool.StringToInt(parent_id)

	if err := request.Verify(menu, request.MenuAddChildVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.CreateChildMenu(menu.ParentId, menu, "hidden"); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 更新菜单
// @Description  更新菜单
// @Router /admin/updateMenu/:id  [post]
func UpdateMenu(c *gin.Context) {
	var menu models.SysMenu

	c.ShouldBind(&menu)

	id := c.Param("id")
	menu.ID = tool.StringToInt(id)

	if err := request.Verify(menu, request.MenuUpVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.UpdateMenu(menu.ID, menu, "id", "parent_id"); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 删除菜单
// @Description  删除菜单
// @Router /admin/deleteMenu/:id  [post]
func DeleteMenu(c *gin.Context) {
	var menu models.SysMenu
	c.ShouldBindUri(&menu)
	if err := request.Verify(menu, request.MenuDelVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.DelMenu(menu); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}
