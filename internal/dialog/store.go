package dialog

type SessionStore struct {
	store map[string]Session
}

func (this *SessionStore) Begin(id string) {
	session, ok := this.session(id)
	if ok {
		session.Abort()
	} else {
		session = *newSession()
	}
}

func (this *SessionStore) Has(id string) bool {
	_, exists := this.session(id)
	return exists
}

func (this *SessionStore) session(id string) (Session, bool) {
	s, ok := this.store[id]
	return s, ok
}

func NewSessionStore() SessionStore {
	return SessionStore{make(map[string]Session, 0)}
}
