package controller

import (
	"github.com/gin-gonic/gin"

	"go-api/app/request"
	"go-api/app/services"
	"go-api/tool"
)

type user struct {
	Name      string `json:"name" xml:"name" form:"name" binding:"required"`
	CreatedBy string `json:"created_by" xml:"created_by" form:"created_by" binding:"lowercase"`
}
type userId struct {
	ID int `uri:"id" binding:"required"`
}

var UserService services.UserContract

func init() {
	UserService = &services.UserService{}
}

// @Summary 用户列表
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

	ret := UserService.GetUserList(maps, tool.DefaultGetOffset(c), 10)
	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	}
	tool.JSONP(c, 0, "查询成功", ret["data"])
}

func GetUser(c *gin.Context) {
	user := new(userId)
	if err := c.ShouldBindUri(user); err != nil {
		tool.JSONP(c, 40001, request.GetError(err), nil)
		return
	}

	ret := UserService.GetUserByID(user.ID)

	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	}
	tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
}

func AddUser(c *gin.Context) {

	user := new(user)
	if err := c.ShouldBind(user); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	maps := make(map[string]interface{})
	maps["Name"] = user.Name

	data := make(map[string]interface{})
	data["Name"] = user.Name
	data["CreatedBy"] = user.CreatedBy
	ret := UserService.AddUser(maps, data)

	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	}
	tool.JSONP(c, 0, ret.GetMsg(), ret["data"])

}

func EditUser(c *gin.Context) {
	userid := new(userId)
	if err := c.ShouldBindUri(userid); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	user := new(user)
	if err := c.ShouldBind(user); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	data := make(map[string]interface{})
	data["name"] = user.Name
	data["created_by"] = user.CreatedBy

	ret := UserService.EditUser(userid.ID, data)

	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), nil)
		return
	}
	tool.JSONP(c, 0, ret.GetMsg(), nil)
}

func DeleteUser(c *gin.Context) {

	userid := new(userId)
	if err := c.ShouldBindUri(userid); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	maps := make(map[string]interface{})
	maps["id"] = userid.ID

	ret := UserService.DeleteUser(maps)

	if !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), nil)
		return
	}
	tool.JSONP(c, 0, ret.GetMsg(), nil)
}
