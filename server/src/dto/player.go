package dto

import (
	"github.com/JeyXeon/poker-easy/model"
)

type Player struct {
	AccountId int
	Name      string
	Hand      []PlayingCard
	Balance   int64
	Bet       int64
}

func AccountsToPlayers(accounts []*model.Account) []Player {
	result := make([]Player, len(accounts))

	for _, account := range accounts {
		result = append(result, Player{
			AccountId: account.ID,
			Name:      account.Username,
			Hand:      make([]PlayingCard, 2),
			Balance:   account.MoneyBalance,
			Bet:       0,
		})
	}

	return result
}
