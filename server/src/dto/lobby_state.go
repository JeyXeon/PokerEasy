package dto

import (
	"errors"
	"github.com/JeyXeon/poker-easy/model"
)

type LobbyState struct {
	Lobby                      *model.Lobby `json:"lobby"`
	places                     []int
	ConnectedPlayersByPlaces   map[int]*Player `json:"playersByPlaces"`
	PlacesByConnectedPlayerIds map[int]int
	GameState                  *GameState
}

func NewLobbyState(lobby *model.Lobby) *LobbyState {
	lobbyState := new(LobbyState)
	lobbyState.Lobby = lobby
	lobbyState.GameState = nil

	maxPlayersAmount := lobby.MaxPlayers

	lobbyState.places = make([]int, maxPlayersAmount)
	for i := 0; i < maxPlayersAmount; i++ {
		lobbyState.places[i] = i
	}

	lobbyState.ConnectedPlayersByPlaces = make(map[int]*Player, maxPlayersAmount)
	for _, place := range lobbyState.places {
		lobbyState.ConnectedPlayersByPlaces[place] = nil
	}

	lobbyState.PlacesByConnectedPlayerIds = make(map[int]int, maxPlayersAmount)

	return lobbyState
}

func (lobbyState *LobbyState) AddPlayer(player *Player) error {
	for _, position := range lobbyState.places {
		reservation := lobbyState.ConnectedPlayersByPlaces[position]
		if reservation == nil {
			lobbyState.ConnectedPlayersByPlaces[position] = player
			lobbyState.PlacesByConnectedPlayerIds[player.ID] = position
			return nil
		}
	}

	return errors.New("all places reserved")
}

func (lobbyState *LobbyState) RemovePlayer(playerId int) error {
	for _, position := range lobbyState.places {
		positionReservation, reserved := lobbyState.ConnectedPlayersByPlaces[position]
		if reserved && positionReservation.ID == playerId {
			lobbyState.ConnectedPlayersByPlaces[position] = nil
			delete(lobbyState.PlacesByConnectedPlayerIds, playerId)
			return nil
		}
	}

	return errors.New("all places reserved")
}

func (lobbyState *LobbyState) ChangeReadyState(playerId int) bool {
	place := lobbyState.PlacesByConnectedPlayerIds[playerId]
	player := lobbyState.ConnectedPlayersByPlaces[place]
	player.IsReady = !player.IsReady

	allReady := true
	for _, player := range lobbyState.ConnectedPlayersByPlaces {
		if player != nil && !player.IsReady {
			allReady = false
			break
		}
	}

	return allReady
}
