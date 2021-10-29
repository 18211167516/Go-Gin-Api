package routes

import (
	"go-api/app/middleware"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

type root struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func testRoute(r *gin.Engine) {

	test := r.Group("/test", middleware.DefaultLog(), middleware.Recovery())
	{

		test.GET("/ping", func(c *gin.Context) {
			tool.NewSecureCookie(c).SetCookie("aaa", "12321323", 86400, "/", "", false, true)
			c.JSON(200, gin.H{
				"message": "",
			})
		})

		test.GET("/panic", func(c *gin.Context) {
			panic("panic")
		})

		test.GET("/someXML", func(c *gin.Context) {
			c.Writer.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>`))
			c.XML(200, root{ID: 1, Name: "baibai", Age: 16})
		})
	}

}
