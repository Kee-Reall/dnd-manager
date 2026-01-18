package domain

type BidStatus byte

const (
	PendingBid BidStatus = iota
	ApprovedBid
	RejectedBid
	CanceledBidBid
)

type Bid struct {
	Id         string
	UserId     string
	CampaignId string
	Status     BidStatus
	TEXT       string
}
