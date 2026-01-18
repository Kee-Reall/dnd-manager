package domain

import "time"

type EditionName string

const (
	//Игра Соотвествует редакции полность, или вносит изменения, но использует значительную базу редакции из Player Handbook
	E5 EditionName = "5e"
	E4 EditionName = "4e"
	// полный кастом. Игра может основываться на какой то из Редакций, но правила изменены колосально
	// Такая игра не ссылается на общие ресурсы
	Homebrew EditionName = "homebrew"
)

type Game struct {
	Date        time.Time  `json:"date"`
	Status      GameStatus `json:"status"`
	Campaign    *Campaign  `json:"-"`
	Description string     `json:"message"`
}

type Edition struct {
	Name EditionName
	Id   string
}

type GameStatus byte

const (
	PlanedGame byte = iota
	InProgressGame
	FinishedGame
	CanceledGame
)
