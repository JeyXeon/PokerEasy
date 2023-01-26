package dto

import (
	"github.com/JeyXeon/poker-easy/model"
)

type Player struct {
	ID           int    `json:"accountId"`
	Username     string `json:"username"`
	MoneyBalance int64  `json:"moneyBalance"`
	IsReady      bool   `json:"isReady"`
}

func AccountToPlayer(account *model.Account) *Player {
	player := new(Player)
	player.ID = account.ID
	player.Username = account.Username
	player.MoneyBalance = account.MoneyBalance
	player.IsReady = false

	return player
}
