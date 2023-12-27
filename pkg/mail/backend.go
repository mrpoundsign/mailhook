package mail

import "github.com/emersion/go-smtp"

func NewBackend(user, pass string) *Backend {
	return &Backend{
		authUser: user,
		authPass: pass,
	}
}

// The Backend implements SMTP server methods.
type Backend struct {
	hooks    []Hook
	authUser string
	authPass string
}

// NewSession is called after client greeting (EHLO, HELO).
func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{
		hooks:    bkd.hooks,
		authUser: bkd.authUser,
		authPass: bkd.authPass,
	}, nil
}

func (b *Backend) AddHook(hook Hook) {
	b.hooks = append(b.hooks, hook)
}
