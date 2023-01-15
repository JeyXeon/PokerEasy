package repository

import (
	"context"
	"github.com/JeyXeon/poker-easy/config"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LobbyRepository struct {
	db *pgxpool.Pool
}

func GetLobbyRepository() *LobbyRepository {
	dbCon := config.GetDbConnection()

	lobbyRepository := new(LobbyRepository)
	lobbyRepository.db = dbCon
	return lobbyRepository
}

func (lobbyRepository *LobbyRepository) CreateLobby(lobby model.Lobby) (*model.Lobby, error) {
	db := lobbyRepository.db

	row := db.QueryRow(
		context.Background(),
		insertLobbyQuery,
		lobby.LobbyName, lobby.MaxPlayers, lobby.CreatorId)
	var createdLobby model.Lobby
	err := row.Scan(&createdLobby)

	return &createdLobby, err
}

func (lobbyRepository *LobbyRepository) GetLobbyById(lobbyId int) (*model.Lobby, error) {
	db := lobbyRepository.db

	var lobby model.Lobby
	err := pgxscan.Get(context.Background(), db, &lobby, getLobbyByIdQuery, lobbyId)
	return &lobby, err
}

func (lobbyRepository *LobbyRepository) GetAllLobbies() (model.Lobbies, error) {
	db := lobbyRepository.db

	var lobbies []model.Lobby
	err := pgxscan.Select(context.Background(), db, &lobbies, getAllLobbiesQuery)
	return lobbies, err
}
