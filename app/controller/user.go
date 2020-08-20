package controller

import (
	"github.com/gin-gonic/gin"

	"go-api/app/models"
	"go-api/app/request"
	"go-api/tool"
)

type user struct {
	Name      string `json:"name" xml:"name" form:"name" binding:"required"`
	CreatedBy string `json:"created_by" xml:"created_by" form:"created_by" binding:"lowercase"`
}
type userId struct {
	ID int `uri:"id" binding:"required"`
}

func GetUsers(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	data["total"] = models.GetUserTotal(maps)
	data["list"] = models.GetUsers(tool.DefaultGetOffset(c), 10, maps)
	tool.JSONP(c, 0, "查询成功", data)
}

func GetUser(c *gin.Context) {
	user := new(userId)
	if err := c.ShouldBindUri(user); err != nil {
		tool.JSONP(c, 40001, request.GetError(err), nil)
		return
	}
	res, err := models.GetUser(user.ID)
	if err != nil {
		tool.JSONP(c, 40001, "暂无数据", nil)
		return
	}
	tool.JSONP(c, 0, "查询成功", res)
}

func AddUser(c *gin.Context) {

	user := new(user)
	if err := c.ShouldBind(user); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	maps := make(map[string]interface{})
	maps["Name"] = user.Name

	if models.ExistUserByMaps(maps) {
		tool.JSONP(c, 40001, "该名称已存在", maps)
		return
	}

	data := make(map[string]interface{})
	data["Name"] = user.Name
	data["CreatedBy"] = user.CreatedBy

	res := models.AddUser(data)
	if !res {
		tool.JSONP(c, 40001, "创建失败", data)
		return
	}
	tool.JSONP(c, 0, "创建成功", nil)
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
	//先查id是否有记录
	if !models.ExistTagByID(userid.ID) {
		tool.JSONP(c, 40001, "ID记录不存在", nil)
		return
	}
	res, err := models.EditUser(userid.ID, data)
	if !res {
		tool.JSONP(c, 40001, "编辑失败", err)
		return
	}
	tool.JSONP(c, 0, "编辑成功", nil)
}

func DeleteUser(c *gin.Context) {

	userid := new(userId)
	if err := c.ShouldBindUri(userid); err != nil {
		tool.JSONP(c, 400, request.GetError(err), nil)
		return
	}

	if !models.ExistTagByID(userid.ID) {
		tool.JSONP(c, 40001, "ID记录不存在", nil)
		return
	}

	maps := make(map[string]interface{})
	maps["id"] = userid.ID

	res, _ := models.DeleteUser(maps)
	if !res {
		tool.JSONP(c, 40001, "删除成功", nil)
		return
	}

	tool.JSONP(c, 0, "删除成功", nil)
}
