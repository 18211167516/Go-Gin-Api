package routes

import (
	"github.com/gin-gonic/gin"
)

func TestRoute(){
	R.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}