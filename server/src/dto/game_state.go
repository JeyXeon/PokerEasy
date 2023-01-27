package dto

import "math/rand"

type GameState struct {
	CurrentRound int
	Deck         []PlayingCard `json:"-"`
	CardsOnTable []PlayingCard
	CurrentTurn  int
	CurrentBet   int64
	Bank         int64
	BigBlind     int
	SmallBlind   int
}

func NewGameState() *GameState {
	gameState := new(GameState)
	gameState.CurrentRound = 0
	gameState.Deck = make([]PlayingCard, len(AvailableCards), len(AvailableCards))
	gameState.CardsOnTable = make([]PlayingCard, 0, 5)
	gameState.CurrentTurn = 0
	gameState.CurrentBet = 0
	gameState.Bank = 0

	copy(gameState.Deck, AvailableCards)
	rand.Shuffle(len(gameState.Deck), func(i, j int) {
		gameState.Deck[i], gameState.Deck[j] = gameState.Deck[j], gameState.Deck[i]
	})

	return gameState
}
