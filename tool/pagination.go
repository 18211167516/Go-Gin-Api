package tool

import (
	"github.com/gin-gonic/gin"

	"go-api/global"
)

//get Offset  limit Optional
func GetOffset(c *gin.Context, limit int) int {
	page := StringToInt(c.DefaultQuery("page", "0"))
	return getOffset(page, limit)
}

//get Offset limit default
func DefaultGetOffset(c *gin.Context) int {
	page := StringToInt(c.DefaultQuery("page", "0"))
	return getOffset(page, global.CF.App.PageSize)
}

func getOffset(page int, limit int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * limit
	}
	return result
}
