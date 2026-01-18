package domain

type CampaignStatus byte

const (
	PreparingCampaign byte = iota
	InProgressCampaign
	PausedCampaign
	FinishedCampaign
	CanceledCampaign
)

type Player struct {
	UserId        string `json:"-"`
	Name          string `json:"playerName"`
	CharacterName string `json:"characterName"`
}

type Campaign struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	MasterId    string     `json:"masterId"`
	Edition     Edition    `json:"edition"`
	Players     []Player   `json:"players"`
	Games       []Game     `json:"games"`
	Status      GameStatus `json:"status"`
}
