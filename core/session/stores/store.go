package stores

import (
	"github.com/18211167516/sessions"
)

type Store interface {
	sessions.Store
	Options(sessions.Options)
}
