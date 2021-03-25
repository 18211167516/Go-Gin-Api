package controller

import (
	"github.com/gin-gonic/gin"

	"go-api/app/request"
	"go-api/app/services"
	"go-api/tool"
)

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

	ret := services.GetUserList(maps, tool.DefaultGetOffset(c), 10)
	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	}
	tool.JSONP(c, 0, "查询成功", ret["data"])
}

// @Summary 单个用户
// @Description  获取单个用户
// @Param id path int true "ID"
// @Tags 用户信息
// @Produce  json
// @Failure 400 {object} tool.JSONRET "参数错误"
// @Failure 20001 {object} tool.JSONRET "Token鉴权失败"
// @Failure 20002 {object} tool.JSONRET "Token已超时"
// @Failure 20004 {object} tool.JSONRET "Token错误"
// @Failure 20005 {object} tool.JSONRET "Token参数不能为空"
// @Success 0 {object} models.UserSwagger "查询成功"
// @Router /api/v1/user/{id} [get]
func GetUser(c *gin.Context) {
	user := new(request.UserId)
	if err := c.ShouldBindUri(user); err != nil {
		tool.JSONP(c, 40001, request.GetError(err), nil)
		return
	}

	ret := services.GetUserByID(user.ID)

	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	}
	tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
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

	user := new(request.User)
	if err := c.ShouldBind(user); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	maps := make(map[string]interface{})
	maps["Name"] = user.Name

	data := make(map[string]interface{})
	data["Name"] = user.Name
	data["CreatedBy"] = user.CreatedBy
	ret := services.AddUser(maps, data)

	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	}
	tool.JSONP(c, 0, ret.GetMsg(), ret["data"])

}

func EditUser(c *gin.Context) {
	userid := new(request.UserId)
	if err := c.ShouldBindUri(userid); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	user := new(request.User)
	if err := c.ShouldBind(user); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	data := make(map[string]interface{})
	data["name"] = user.Name
	data["created_by"] = user.CreatedBy

	ret := services.EditUser(userid.ID, data)

	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), nil)
		return
	}
	tool.JSONP(c, 0, ret.GetMsg(), nil)
}

func DeleteUser(c *gin.Context) {

	userid := new(request.UserId)
	if err := c.ShouldBindUri(userid); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	maps := make(map[string]interface{})
	maps["id"] = userid.ID

	ret := services.DeleteUser(maps)

	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), nil)
		return
	}
	tool.JSONP(c, 0, ret.GetMsg(), nil)
}
