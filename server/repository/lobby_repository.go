package repository

import (
	"github.com/JeyXeon/poker-easy/config"
	"github.com/JeyXeon/poker-easy/model"
	"gorm.io/gorm"
)

type LobbyRepository struct {
	db *gorm.DB
}

func GetLobbyRepository() *LobbyRepository {
	dbCon := config.GetDbConnection()
	return &LobbyRepository{db: dbCon}
}

func (lobbyRepository *LobbyRepository) CreateLobby(lobby model.Lobby) (model.Lobby, error) {
	db := lobbyRepository.db

	tx := db.Begin()
	if err := tx.Create(&lobby).Error; err != nil {
		tx.Rollback()
		return model.Lobby{}, err
	}
	tx.First(&lobby).Where("creator_id = ?", lobby.CreatorId)
	tx.Commit()

	return lobby, nil
}

func (lobbyRepository *LobbyRepository) GetLobbyById(lobbyId int) model.Lobby {
	db := lobbyRepository.db

	lobby := model.Lobby{}
	db.First(&lobby, "lobby_id = ?", lobbyId)
	return lobby
}

func (lobbyRepository *LobbyRepository) GetAllLobbies() model.Lobbies {
	db := lobbyRepository.db

	lobbies := make(model.Lobbies, 0)
	db.Find(&lobbies)

	return lobbies
}
