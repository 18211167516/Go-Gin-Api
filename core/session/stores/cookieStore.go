package stores

import (
	"github.com/18211167516/sessions"
)


type cookieStore struct {
	*sessions.CookieStore
}

func NewCookieStore(keyPairs ...[]byte) Store {
	return &cookieStore{sessions.NewCookieStore(keyPairs...)}
}

func (c *cookieStore) Options(options sessions.Options) {
	c.CookieStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
