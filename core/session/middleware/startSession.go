package middleware

import (
	"go-api/core/session"

	"github.com/gin-gonic/gin"
)

func StartSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		cursession := session.NewSession(c, nil)
		c.Set(cursession.Name(), cursession)
		c.Next()
	}
}
