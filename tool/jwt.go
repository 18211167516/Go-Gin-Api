package tool

import (
	jwt "github.com/golang-jwt/jwt"

	"go-api/global"
)

var JwtSecret = []byte(global.CF.App.JwtSecret)

type Claims struct {
	ID             string //用户ID
	Name           string //用户名称
	Type           int    //类型
	RuleName       string //角色名称 可能存在多个角色名
	StandardClaims jwt.StandardClaims
}

/*实现接口*/
func (c Claims) Valid() error {
	return c.StandardClaims.Valid()
}

/*获取自定义载荷用户ID*/
func (c Claims) GetID() string {
	return c.ID
}

/*获取自定义载荷用户名称*/
func (c Claims) GetName() string {
	return c.Name
}

/*获取自定义载荷用户角色名称*/
func (c Claims) GetRuleName() string {
	return c.RuleName
}

/*生成jwt-token*/
func GenerateToken(claims Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.GetSigningMethod(global.CF.App.SigningMethod), claims)

	return token.SignedString(JwtSecret)
}

/*简析jwt-token*/
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
