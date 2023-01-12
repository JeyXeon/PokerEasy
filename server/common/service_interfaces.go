package common

import (
	"github.com/JeyXeon/poker-easy/model"
	"github.com/gofiber/websocket/v2"
)

type AccountService interface {
	SaveNewAccount(accountDto model.Account) (*model.Account, error)
	GetAccountById(accountId int) (*model.Account, error)
	UpdateAccount(account *model.Account)
	RemoveLobbyConnection(accountId int)
}

type LobbyService interface {
	SaveNewLobby(lobby model.Lobby) (*model.Lobby, error)
	GetLobbyById(lobbyId int) (*model.Lobby, error)
	GetAllLobbies() (model.Lobbies, error)
}

type GameService interface {
	ListenWebsocket(conn *websocket.Conn)
}
