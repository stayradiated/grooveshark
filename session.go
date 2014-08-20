package grooveshark

import "time"

type SessionToken string

type Session struct {
	Salt           string
	Client         string
	ClientRevision string
	Token          SessionToken
	Country        Country
	SessionId      SessionId
	LastUpdated    time.Time
}

func NewSession() (session *Session) {
	return session
}

func (s *Session) Connect() {
}

func (s *Session) Check() {
}

func (s *Session) GetToken() SessionToken {
	s.Check()
	return s.Token
}
