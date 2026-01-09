package domain

type Status byte

const (
	Pending Status = iota
	Approved
	Rejected
)

type Bid struct {
	Id       string
	PlayerId string
	GameId   string
	Status   Status
}
