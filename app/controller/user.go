package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/app/models"
)

func GetUsers(c *gin.Context) {
	maps := make(map[string]interface{})
    data := make(map[string]interface{})
	data["total"] = models.GetUserTotal(maps)
	data["list"]  = models.GetUsers(0,10,maps)
	c.JSON(200, gin.H{
		"status":  "200",
		"message": "查询成功",
		"data":    data,
	})
}

func GetUser(c *gin.Context) {

}

func AddUser(c *gin.Context) {

}

func EditUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context){

}