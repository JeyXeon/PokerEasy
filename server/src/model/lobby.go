package model

type Lobby struct {
	ID         int    `json:"lobbyId" db:"lobby_id"`
	LobbyName  string `json:"lobbyName" db:"lobby_name"`
	MaxPlayers int    `json:"maxPlayers" db:"players_amount"`
	CreatorId  int    `json:"creatorId" db:"creator_id"`
	players    []Account
}

type Lobbies []Lobby
