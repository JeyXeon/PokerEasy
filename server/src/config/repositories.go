package config

import (
	"github.com/JeyXeon/poker-easy/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppRepositories struct {
	accountRepository *repository.AccountRepository
	lobbyRepository   *repository.LobbyRepository
}

func GetAppRepositories(db *pgxpool.Pool) *AppRepositories {
	appRepositories := new(AppRepositories)
	appRepositories.accountRepository = repository.GetAccountRepository(db)
	appRepositories.lobbyRepository = repository.GetLobbyRepository(db)

	return appRepositories
}
