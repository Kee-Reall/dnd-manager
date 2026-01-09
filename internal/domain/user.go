package domain

type Maker string
type Role string

const (
	TG     Maker = "tg"
	Player Role  = "player"
	Master Role  = "master"
	Admin  Role  = "admin"
)

type UserViewId struct {
	// Внешний Айди из Вью Системы - В данном случае телеграмма
	ID string
	//Метка Вью системы, пока только телеграм
	Marker Maker
}

type User struct {
	ID     string     `json:"-"`
	ViewId UserViewId `json:"-"`
	Name   string     `json:"name"`
	Role   Role       `json:"role"`
}
