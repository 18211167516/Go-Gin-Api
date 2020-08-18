package routes

import (
	"github.com/gin-gonic/gin"
)

type root struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func testRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.Writer.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>`))
		c.XML(200, root{ID: 1, Name: "baibai", Age: 16})
	})
}
