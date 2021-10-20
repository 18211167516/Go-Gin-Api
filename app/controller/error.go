package controller

import (
	"github.com/gin-gonic/gin"

	"go-api/tool"
)

// @Summary 错误页
// @Description  错误页
// @Router /admin/error/:code/:message [get]
func Error(c *gin.Context) {
	code := c.Param("code")
	message := c.Param("message")
	tool.ViewErr(c, tool.StringToInt(code), message)
}
