package common

import "github.com/JeyXeon/poker-easy/model"

type AccountRepository interface {
	CreateAccount(account model.Account) (*model.Account, error)
	GetAccountById(accountId int) (*model.Account, error)
	UpdateAccount(account *model.Account)
	RemoveLobbyConnection(accountId int)
}

type LobbyRepository interface {
	CreateLobby(account model.Lobby) (*model.Lobby, error)
	GetLobbyById(accountId int) (*model.Lobby, error)
	GetAllLobbies() (model.Lobbies, error)
}
