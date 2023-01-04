package model

import "fmt"
import "strings"

type Lobby struct {
	ID         int `gorm:"column:lobby_id"`
	LobbyName  string
	MaxPlayers int `gorm:"column:players_amount"`
	CreatorId  int
	Players    []Account `gorm:"foreignKey:ConnectedLobbyId"`
}

func (Lobby) TableName() string {
	return "lobby"
}

func (lobby Lobby) ToString() string {
	return fmt.Sprintf("{ID: %d, LobbyName: %s, MaxPlayers: %d}", lobby.ID, lobby.LobbyName, lobby.MaxPlayers)
}

type Lobbies []Lobby

func (lobbies Lobbies) ToString() string {
	sb := strings.Builder{}
	for _, lobby := range lobbies {
		sb.WriteString(lobby.ToString())
	}
	return sb.String()
}
