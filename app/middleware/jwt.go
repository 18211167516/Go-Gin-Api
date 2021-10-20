package middleware

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"go-api/tool"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code  = tool.SUCCESS
			token string
		)

		if s, ok := c.GetQuery("token"); ok {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			code = tool.ERROR_AUTH_CHECK_TOKEN_EMPTY
		} else {
			claims, err := tool.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = tool.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = tool.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}

			c.Set("claims", claims)
		}

		if gin.Mode() == "debug" {
			code = tool.SUCCESS
		}
		if code != tool.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  tool.GetMsg(code),
				"data": c.Request.URL.RequestURI(),
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
