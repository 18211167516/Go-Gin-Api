package tool

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"go-api/config"
)

var (
	JwtSecret     = []byte(config.AppSetting.JwtSecret)
	JwtExpiresAt  = config.AppSetting.JwtExpiresAt
	SigningMethod = config.AppSetting.SigningMethod
)

type Claims struct {
	Appkey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GenerateToken(appkey, app_secret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(JwtExpiresAt)

	claims := Claims{
		appkey,     //可以考虑加密
		app_secret, //可以考虑加密
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-api",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.GetSigningMethod(SigningMethod), claims)
	token, err := tokenClaims.SignedString(JwtSecret)

	return token, err
}

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
