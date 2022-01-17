package controller

import (
	"go-api/app/request"
	"go-api/app/services"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

// @Summary 任务列表视图
// @Description  任务列表视图
// @Router /admin/CronListView [get]
func CronListView(c *gin.Context) {
	uid := c.GetString("uid")
	view_route := c.Request.URL.RequestURI()
	data := tool.M{
		"dataUrl":    "/admin/getCrons",
		"dataMethod": "POST",
		"runUrl":     services.GetButtonPermission(uid, view_route, "/admin/runCron/:id"),
		"delUrl":     services.GetButtonPermission(uid, view_route, "/admin/deleteCron/:id"),
	}
	tool.HTML(c, "cron/cron_list.html", data)
}

// @Summary 任务列表
// @Description  任务列表
// @Router /admin/getCrons [post]
func GetCrons(c *gin.Context) {
	ret := services.GetCronList()
	tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
}

func RunCron(c *gin.Context) {
	id := c.Param("id")
	ids := struct {
		id string
	}{id: id}

	if err := request.Verify(ids, request.DelVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.RunCron(id); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 删除任务
// @Description  删除任务
// @Router /admin/deleteCron/:id [post]
func DeleteCron(c *gin.Context) {
	id := c.Param("id")
	ids := struct {
		id string
	}{id: id}

	if err := request.Verify(ids, request.DelVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	if ret := services.DelCron(id); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}
