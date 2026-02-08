package domain

type MakerTag string
type Role byte

const (
	NoRole Role = iota
	PlayerRole
	MasterRole
	AdminRole
	TG MakerTag = "tg"
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

func AllRole() []Role {
	return []Role{NoRole, PlayerRole, MasterRole, AdminRole}
}

func MasterFunctionsRole() []Role {
	return []Role{MasterRole, AdminRole}
}

func MaxRole() Role {
	return AdminRole
}

type User struct {
	ID     string     `json:"-"`
	Marker UserMarker `json:"-"`
	Name   string     `json:"name"`
	Role   Role       `json:"role"`
}
