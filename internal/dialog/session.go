package dialog

type Session struct {
	cap int
	len int
}

func (s *Session) Abort() {
}

func (s *Session) Len() int {
	return s.len
}

func (s *Session) Cap() int {
	return s.cap
}

func newSession() *Session {
	return &Session{}
}
