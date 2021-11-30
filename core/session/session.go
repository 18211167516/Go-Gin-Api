package session

import (
	"log"
	"net/http"

	"go-api/core/session/stores"
	"go-api/global"

	"github.com/18211167516/sessions"
	"github.com/gin-gonic/gin"
)

var DefaultKey = "SESSION"
var errorFormat = "[sessions] ERROR! %s\n"

type Session interface {
	// Get returns the session value associated to the given key.
	Get(key interface{}) interface{}
	// Set sets the session value associated to the given key.
	Set(key interface{}, val interface{})
	// Delete removes the session value associated to the given key.
	Delete(key interface{})
	// Clear deletes all values in the session.
	Clear()
	Options(*sessions.Options)
	// AddFlash adds a flash message to the session.
	// A single variadic argument is accepted, and it is optional: it defines the flash key.
	// If not defined "_flash" is used by default.
	AddFlash(value interface{}, vars ...string)
	// Flashes returns a slice of flash messages from the session.
	// A single variadic argument is accepted, and it is optional: it defines the flash key.
	// If not defined "_flash" is used by default.
	Flashes(vars ...string) []interface{}
	// Save saves all sessions used during the current request.
	Save() error
}

type session struct {
	name    string
	request *http.Request
	store   sessions.Store
	session *sessions.Session
	written bool
	writer  http.ResponseWriter
}

func NewSession(c *gin.Context, store stores.Store) *session {
	if store == nil {
		switch global.VP.GetString("session.driver") {
		case "cookie":
			store = stores.NewCookieStore([]byte(global.VP.GetString("session.cookie")))
		case "redis":
			store = stores.GetRedisStore()
		default:
			store = stores.NewCookieStore([]byte(global.VP.GetString("session.cookie")))
		}
	}

	return &session{
		name:    global.VP.GetString("session.cookie"),
		request: c.Request,
		store:   store,
		writer:  c.Writer,
	}
}

func (s *session) Options(options *sessions.Options) {
	s.session.Options = options
}

func (s *session) Name() string {
	return s.name
}

func (s *session) Get(key interface{}) interface{} {
	return s.Session().Values[key]
}

func (s *session) Set(key interface{}, val interface{}) {
	s.Session().Values[key] = val
	s.written = true
}

func (s *session) Delete(key interface{}) {
	delete(s.Session().Values, key)
	s.written = true
}

func (s *session) Clear() {
	for key := range s.Session().Values {
		s.Delete(key)
	}
}

func (s *session) AddFlash(value interface{}, vars ...string) {
	s.Session().AddFlash(value, vars...)
	s.written = true
}

func (s *session) Flashes(vars ...string) []interface{} {
	s.written = true
	return s.Session().Flashes(vars...)
}

func (s *session) Save() error {
	if s.Written() {
		e := s.Session().Save(s.request, s.writer)
		if e == nil {
			s.written = false
		}
		return e
	}
	return nil
}

//获取session对象
func (s *session) Session() *sessions.Session {
	if s.session == nil {
		var err error
		s.session, err = s.store.Get(s.request, s.name)
		if err != nil {
			log.Printf(errorFormat, err)
		}

		options := &sessions.Options{
			Path:     global.VP.GetString("session.path"),
			Domain:   global.VP.GetString("session.domain"),
			MaxAge:   global.VP.GetInt("session.maxAge"),
			Secure:   global.VP.GetBool("session.secure"),
			HttpOnly: global.VP.GetBool("session.http_only"),
			SameSite: http.SameSite(global.VP.GetInt("session.same_site")),
		}
		s.Options(options)
	}
	return s.session
}

func (s *session) Written() bool {
	return s.written
}

func Default(c *gin.Context) Session {
	return c.MustGet(global.VP.GetString("session.cookie")).(Session)
}
