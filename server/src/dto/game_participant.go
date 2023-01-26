package dto

type GameParticipant struct {
	Player *Player
	Hand   []PlayingCard `json:"hand"`
	Bet    int64         `json:"bet"`
}
