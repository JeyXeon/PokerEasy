package service

import (
	"github.com/JeyXeon/poker-easy/model"
	"github.com/JeyXeon/poker-easy/repository"
)

type LobbyRepository interface {
	CreateLobby(account model.Lobby) (*model.Lobby, error)
	GetLobbyById(accountId int) *model.Lobby
	GetAllLobbies() model.Lobbies
}

type LobbyService struct {
	lobbyRepository LobbyRepository
}

func GetLobbyService() *LobbyService {
	lobbyRepository := repository.GetLobbyRepository()
	return &LobbyService{lobbyRepository: lobbyRepository}
}

func (lobbyService *LobbyService) SaveNewLobby(lobbyDto model.Lobby) *model.Lobby {
	lobbyRepository := lobbyService.lobbyRepository
	createdLobby, err := lobbyRepository.CreateLobby(lobbyDto)
	if err != nil {

	}
	return createdLobby
}

func (lobbyService *LobbyService) GetLobbyById(lobbyId int) *model.Lobby {
	lobbyRepository := lobbyService.lobbyRepository
	existingLobby := lobbyRepository.GetLobbyById(lobbyId)
	return existingLobby
}

func (lobbyService *LobbyService) GetAllLobbies() model.Lobbies {
	lobbyRepository := lobbyService.lobbyRepository
	existingLobbies := lobbyRepository.GetAllLobbies()
	return existingLobbies
}
