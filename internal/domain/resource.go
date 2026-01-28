package domain

type ResourceScope byte

const (
	ScopeEdition ResourceScope = iota
	ScopeCampaign
)

type Resource struct {
	ID      string
	MediaID string
	Title   string
	Scope   ResourceScope
	OwnerID string // game_id или edition_id
}
