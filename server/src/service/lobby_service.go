package service

import (
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/model"
)

type LobbyService struct {
	lobbyRepository common.LobbyRepository
}

func GetLobbyService(lobbyRepository common.LobbyRepository) *LobbyService {
	lobbyService := new(LobbyService)
	lobbyService.lobbyRepository = lobbyRepository
	return lobbyService
}

func (lobbyService *LobbyService) SaveNewLobby(lobbyDto model.Lobby) (*model.Lobby, error) {
	lobbyRepository := lobbyService.lobbyRepository
	createdLobby, err := lobbyRepository.CreateLobby(lobbyDto)

	return createdLobby, err
}

func (lobbyService *LobbyService) GetLobbyById(lobbyId int) (*model.Lobby, error) {
	lobbyRepository := lobbyService.lobbyRepository
	existingLobby, err := lobbyRepository.GetLobbyById(lobbyId)
	return existingLobby, err
}

func (lobbyService *LobbyService) GetAllLobbies() (model.Lobbies, error) {
	lobbyRepository := lobbyService.lobbyRepository
	existingLobbies, err := lobbyRepository.GetAllLobbies()
	return existingLobbies, err
}
