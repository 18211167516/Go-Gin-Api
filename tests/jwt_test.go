package tests

import (
	"fmt"
	"go-api/tool"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestJwt(t *testing.T) {
	claims := tool.Claims{
		ID:       "1",
		Name:     "白葱花",
		Type:     1,
		RuleName: "超级管理员",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*7, //过期时间
			Issuer:    "go-api",                       //签发人
		},
	}
	token, err := tool.GenerateToken(claims)
	fmt.Println(token, err)

	fmt.Println(tool.ParseToken(token))
}
