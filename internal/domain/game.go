package domain

type Edition string

const (
	//Игра Соотвествует редакции полность, или вносит незначительные изменения
	E5 Edition = "5e"
	E4 Edition = "4e"
	//эти редакции означают что игра в целом соотвествует одной из Редакций, но есть значительные изменения
	E5hb Edition = "e5hb"
	E4hb Edition = "e4hb"
	// полный кастом. Игра может основываться на какой то из Редакций, но правила изменены колосально
	Homebrew Edition = "homebrew"
)

type Game struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	MasterId    string   `json:"masterId"`
	Edition     string   `json:"edition"`
	Players     []string `json:"players"`
}
