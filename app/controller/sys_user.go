package controller

import (
	"github.com/gin-gonic/gin"

	"go-api/app/models"
	"go-api/app/request"
	"go-api/app/services"
	"go-api/tool"
)

// @Summary 用户列表
// @Description  用户列表
// @Router /admin/usersView [get]
func UsersView(c *gin.Context) {
	uid := c.GetString("uid")
	view_route := c.Request.URL.RequestURI()
	data := tool.M{
		"dataUrl":    "/admin/getUsers",
		"dataMethod": "POST",
		"addUrl":     services.GetButtonPermission(uid, view_route, "/admin/createUser"),
		"editUrl":    services.GetButtonPermission(uid, view_route, "/admin/updateUser/:id"),
		"delUrl":     services.GetButtonPermission(uid, view_route, "/admin/deleteUser/:id"),
		"ruleUrl":    services.GetButtonPermission(uid, view_route, "/admin/userRuleView/:id"),
		"changeUrl":  services.GetButtonPermission(uid, view_route, "/admin/changePassword/:id"),
	}

	tool.HTML(c, "user/user_list.html", data)
}

// @Summary 个人信息页
// @Description  个人信息页
// @Router /admin/userView/ [get]
func UserView(c *gin.Context) {
	id := c.GetString("uid")
	ret := services.GetUserByID(id)

	data := tool.M{
		"data":      ret["data"],
		"UpdateUrl": "/admin/changeOwnInfo",
	}

	tool.HTML(c, "user/user.html", data)
}

// @Summary 修改个人信息
// @Description  修改个人信息
// @Router /admin/changeOwnInfo [post]
func ChangeOwnInfo(c *gin.Context) {
	var user models.SysUser

	c.ShouldBind(&user)

	id := c.GetString("uid")
	user.ID = tool.StringToInt(id)

	if err := request.Verify(user, request.UserChangeOwnVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.UpdateUser(user.ID, user, "id", "password", "type", "status"); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 修改密码
// @Description  修改密码
// @Router /admin/changePassword/:id [post]
func ChangePassword(c *gin.Context) {
	var user models.SysUser

	c.ShouldBind(&user)

	id := c.Param("id")
	user.ID = tool.StringToInt(id)

	if err := request.Verify(user, request.UserChangePasswordVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if err := request.VerifyValue(user.Password, user.ConfirmPassword, "eqfield"); err != nil {
		tool.JSONP(c, 400, "确认密码和新密码不一致", nil)
		return
	}

	if ret := services.ChangePassword(user.ID, user); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 修改自己密码
// @Description  修改自己密码
// @Router /admin/changePassword/:id [post]
func ChangeOwnPassword(c *gin.Context) {
	var user request.ChangeOwnPassword

	c.ShouldBind(&user)

	if err := request.Verify(user, request.UserChangeOwnPasswordVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if err := request.VerifyValue(user.NewPassword, user.ConfirmPassword, "eqfield"); err != nil {
		tool.JSONP(c, 400, "确认密码和新密码不一致", nil)
		return
	}

	//获取当前登录用户ID
	id := c.GetString("uid")
	if ret := services.ChangeOwnPassword(id, user.OldPassword, user.NewPassword); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 用户列表
// @Description  获取用户列表
// @Tags 用户信息
// @Produce  json
// @Failure 400 {object} tool.JSONRET "参数错误"
// @Failure 20001 {object} tool.JSONRET "Token鉴权失败"
// @Failure 20002 {object} tool.JSONRET "Token已超时"
// @Failure 20004 {object} tool.JSONRET "Token错误"
// @Failure 20005 {object} tool.JSONRET "Token参数不能为空"
// @Success 0 {object} models.UserSwagger "查询成功"
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	maps := make(map[string]interface{})
	name := c.PostForm("name")
	Type := c.PostForm("type")
	if name != "" {
		maps["name"] = name
	}
	if Type != "" {
		maps["type"] = Type
	}
	if ret := services.GetUserList(maps, tool.DefaultGetOffset(c), 10); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, "查询成功", ret["data"])
	}
}

// @Summary 新增用户
// @Description  新增用户
// @Accept multipart/form-data*
// @Param name formData string true "Name"
// @Param created_by formData string true "CreatedBy"
// @Tags 用户信息
// @Failure 400 {object} tool.JSONRET "参数错误"
// @Failure 20001 {object} tool.JSONRET "Token鉴权失败"
// @Failure 20002 {object} tool.JSONRET "Token已超时"
// @Failure 20004 {object} tool.JSONRET "Token错误"
// @Failure 20005 {object} tool.JSONRET "Token参数不能为空"
// @Success 0 {object} models.UserSwagger "创建成功"
// @Router /api/v1/users [post]
func AddUser(c *gin.Context) {

	var user models.SysUser
	c.ShouldBind(&user)
	if err := request.Verify(user, request.UserAddVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.Register(user); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 编辑用户
// @Description  编辑用户
// @Router /admin/updateUser/:id [post]
func EditUser(c *gin.Context) {
	var user models.SysUser

	c.ShouldBind(&user)

	id := c.Param("id")
	user.ID = tool.StringToInt(id)

	if err := request.Verify(user, request.UserUpVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.UpdateUser(user.ID, user, "id", "password"); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 删除用户
// @Description  删除用户
// @Router /admin/deleteUser/:id [post]
func DeleteUser(c *gin.Context) {
	var user models.SysUser
	c.ShouldBindUri(&user)
	if err := request.Verify(user, request.UserDelVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.DeleteUser(user); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}
