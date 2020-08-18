package controller

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"go-api/app/models"
	"go-api/tool"
)

func GetUsers(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	data["total"] = models.GetUserTotal(maps)
	data["list"] = models.GetUsers(tool.DefaultGetOffset(c), 10, maps)
	c.JSONP(200, gin.H{"error_code": 0, "msg": tool.GetMsg(0, "查询成功"), "data": data})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	res, err := models.GetUser(tool.StringToInt(id))
	if err != nil {
		c.JSONP(200, gin.H{"error_code": 40001, "msg": tool.GetMsg(40001, "暂无数据"), "err": fmt.Sprint(err)})
	} else {
		c.JSONP(200, gin.H{"error_code": 0, "msg": tool.GetMsg(0, "查询成功"), "data": res})
	}
}

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	created_by := c.PostForm("created_by")
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	if name == "" {
		c.JSONP(200, gin.H{"error_code": 40001, "msg": tool.GetMsg(40001, "名称不能为空")})
	} else {
		data["Name"] = name
		maps["Name"] = name
		data["CreatedBy"] = created_by

		if !models.ExistUserByMaps(maps) {
			res := models.AddUser(data)
			if res {
				c.JSONP(200, gin.H{"error_code": 0, "msg": tool.GetMsg(0, "创建成功"), "data": data})
			} else {
				c.JSONP(200, gin.H{"error_code": 40001, "msg": tool.GetMsg(40001, "创建失败")})
			}
		} else {
			c.JSONP(200, gin.H{"error_code": 40001, "msg": tool.GetMsg(40001, "该名称已存在"), "map": maps})
		}
	}

}

func EditUser(c *gin.Context) {
	id := tool.StringToInt(c.Param("id"))
	name := c.PostForm("name")
	created_by := c.PostForm("created_by")

	if !models.ExistTagByID(id) {
		c.JSONP(200, gin.H{
			"error_code": 40001,
			"msg":        tool.GetMsg(40001, "ID不存在"),
		})
	} else {
		data := make(map[string]interface{})
		data["name"] = name
		data["created_by"] = created_by
		res, err := models.EditUser(id, data)
		if res {
			c.JSONP(200, gin.H{
				"error_code": 0,
				"msg":        tool.GetMsg(40001, "编辑成功"),
			})
		} else {
			c.JSONP(200, gin.H{
				"error_code": 40001,
				"msg":        tool.GetMsg(40001, "编辑失败"),
				"err":        err,
			})
		}
	}
}

func DeleteUser(c *gin.Context) {
	id := tool.StringToInt(c.Param("id"))

	if id <= 0 {
		c.JSONP(200, gin.H{
			"error_code": 40001,
			"msg":        tool.GetMsg(40001, "ID不存在"),
		})
	} else {

		if models.ExistTagByID(id) {
			maps := make(map[string]interface{})

			maps["id"] = id
			res, err := models.DeleteUser(maps)
			if res {
				c.JSONP(200, gin.H{
					"error_code": 0,
					"msg":        tool.GetMsg(40001, "删除成功"),
					"err":        err,
				})
			} else {
				c.JSONP(200, gin.H{
					"error_code": 40001,
					"msg":        tool.GetMsg(40001, "删除失败"),
					"err":        err,
				})
			}
		} else {
			c.JSONP(200, gin.H{
				"error_code": 40001,
				"msg":        tool.GetMsg(40001, "信息不存在"),
			})
		}

	}
}
