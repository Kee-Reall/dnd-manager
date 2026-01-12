package domain

type BidStatus byte

const (
	Pending BidStatus = iota
	Approved
	Rejected
	CanceledBid
)

type Bid struct {
	Id       string
	PlayerId string
	GameId   string
	Status   BidStatus
}
