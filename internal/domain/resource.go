package domain

type ResourceScope byte

const (
	ScopeGlobal ResourceScope = iota
	ScopeGame
)

type Resource struct {
	ID      string
	MediaID string
	Title   string
	Scope   ResourceScope
	OwnerID string // game_id или edition_id
}
