package routes

import (
	"go-api/app/middleware"
	"go-api/app/response"
	"go-api/core/session"
	coremiddleware "go-api/core/session/middleware"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

type root struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func testRoute(r *gin.Engine) {

	test := r.Group("/test", middleware.DefaultLog(), middleware.Recovery(), coremiddleware.StartSession())
	{

		test.GET("/ping", func(c *gin.Context) {
			/* maps := make(map[string]interface{})
			ret := services.GetRoleList(maps, tool.DefaultGetOffset(c), 10) */
			s := session.Default(c)
			count := s.Get("count")
			s.Set("count", 111)
			s.Save()
			tool.JSONP(c, 0, "成功", count)
		})

		test.GET("/panic", func(c *gin.Context) {
			box := response.SysLoginUserResponse{
				ID:       "1",
				Name:     "白",
				RealName: "Bai",
				Type:     2,
				Password: "string",
			}
			s := session.Default(c)
			count := s.Get("user")
			s.Set("user", tool.StructToJson(box))
			s.Save()

			tool.JSONP(c, 0, "层高", count)
		})

		test.GET("/someXML", func(c *gin.Context) {
			c.Writer.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>`))
			c.XML(200, root{ID: 1, Name: "baibai", Age: 16})
		})
	}

}
