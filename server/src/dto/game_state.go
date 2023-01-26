package dto

type GameState struct {
	CurrentRound    int
	PlayersByPlaces map[int]*Player
	Deck            []PlayingCard
	CardsOnTable    []PlayingCard
	CurrentTurn     int
	CurrentBet      int64
	Bank            int64
}

func NewGameState(playersByPlaces map[int]*Player) *GameState {
	gameState := new(GameState)
	gameState.CurrentRound = 0
	gameState.PlayersByPlaces = playersByPlaces
	gameState.Deck = AvailableCards
	gameState.CardsOnTable = make([]PlayingCard, 5)
	gameState.CurrentTurn = 0
	gameState.CurrentBet = 0
	gameState.Bank = 0

	return gameState
}
