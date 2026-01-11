package domain

type MakerTag string
type Role byte

const (
	TG         MakerTag = "tg"
	NoRole     Role     = 0
	PlayerRole Role     = 1
	MasterRole Role     = 2
	AdminRole  Role     = 3
)

func (r Role) String() (roleStr string) {
	switch r {
	case AdminRole:
		roleStr = "Admin"
	case PlayerRole:
		roleStr = "Player"
	case MasterRole:
		roleStr = "Dungeon Master"
	default:
		roleStr = "undefined Role"
	}
	return
}

type User struct {
	ID     string     `json:"-"`
	Marker UserMarker `json:"-"`
	Name   string     `json:"name"`
	Role   Role       `json:"role"`
}
