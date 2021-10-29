package tool

import (
	"go-api/global"
	"sync"

	"github.com/18211167516/encrypt"
	"github.com/gin-gonic/gin"
)

var SecurityCookie *SecureCookie
var once sync.Once

type SecureCookie struct {
	encrypt *encrypt.Encrypt
	c       *gin.Context
}

func NewSecureCookie(c *gin.Context) *SecureCookie {
	once.Do(func() {
		var (
			encryptMode = global.VP.GetString("cookie.Mode")
			key         = global.VP.GetString("cookie.key")
			iv          = global.VP.GetString("cookie.iv")
			padding     = global.VP.GetString("cookie.padding")
		)
		crypt := encrypt.NewEncrypt(encryptMode, []byte(key), []byte(iv), padding)
		SecurityCookie = &SecureCookie{
			encrypt: crypt,
			c:       c,
		}
	})
	return SecurityCookie
}

func (cookie *SecureCookie) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {

	value = cookie.encrypt.EncryptBase64([]byte(value))
	cookie.c.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

func (cookie *SecureCookie) GetCookie(name string) (string, error) {
	value, err := cookie.c.Cookie(name)
	if err != nil {
		return "", err
	}

	if decrypt, err := cookie.encrypt.Base64Decrypt(value); err != nil {
		return "", err
	} else {
		return string(decrypt), nil
	}

}
