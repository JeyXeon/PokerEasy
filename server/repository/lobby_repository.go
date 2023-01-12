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
		`INSERT INTO lobby(lobby_name, players_amount, creator_id) VALUES ($1, $2, $3) RETURNING (lobby_id, lobby_name, players_amount, creator_id)`,
		lobby.LobbyName, lobby.MaxPlayers, lobby.CreatorId)
	var createdLobby model.Lobby
	err := row.Scan(&createdLobby)

	return &createdLobby, err
}

func (lobbyRepository *LobbyRepository) GetLobbyById(lobbyId int) (*model.Lobby, error) {
	db := lobbyRepository.db

	var lobby model.Lobby
	err := pgxscan.Get(context.Background(), db, &lobby, `SELECT * FROM lobby WHERE players_amount = $1;`, lobbyId)
	return &lobby, err
}

func (lobbyRepository *LobbyRepository) GetAllLobbies() (model.Lobbies, error) {
	db := lobbyRepository.db

	var lobbies []model.Lobby
	err := pgxscan.Select(context.Background(), db, &lobbies, "SELECT * FROM lobby")
	return lobbies, err
}
