package common

import (
	"github.com/JeyXeon/poker-easy/model"
	"github.com/gofiber/websocket/v2"
)

type LobbyService interface {
	SaveNewLobby(lobby model.Lobby) *model.Lobby
	GetLobbyById(lobbyId int) *model.Lobby
	GetAllLobbies() model.Lobbies
}

type GameService interface {
	ListenWebsocket(conn *websocket.Conn)
}

type AccountService interface {
	SaveNewAccount(accountDto model.Account) *model.Account
	GetAccountById(accountId int) *model.Account
	UpdateAccount(account *model.Account)
	RemoveLobbyConnection(accountId int)
}
