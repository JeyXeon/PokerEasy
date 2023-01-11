package dto

import "github.com/JeyXeon/poker-easy/model"

type GameState struct {
	CurrentRound int
	Players      []Player
	Deck         []PlayingCard
	CardsOnTable []PlayingCard
	CurrentTurn  int
	CurrentBet   int64
	Bank         int64
}

func NewGameState(accounts []*model.Account) *GameState {
	gameState := new(GameState)
	gameState.CurrentRound = 0
	gameState.Players = AccountsToPlayers(accounts)
	gameState.Deck = AvailableCards
	gameState.CardsOnTable = make([]PlayingCard, 5)
	gameState.CurrentTurn = 0
	gameState.CurrentBet = 0
	gameState.Bank = 0

	return gameState
}
