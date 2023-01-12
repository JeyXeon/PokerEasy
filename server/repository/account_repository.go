package repository

import (
	"context"
	"github.com/JeyXeon/poker-easy/config"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepository struct {
	db *pgxpool.Pool
}

func GetAccountRepository() *AccountRepository {
	dbCon := config.GetDbConnection()
	accountRepository := new(AccountRepository)
	accountRepository.db = dbCon
	return accountRepository
}

func (accountRepository *AccountRepository) CreateAccount(account model.Account) (*model.Account, error) {
	db := accountRepository.db

	var err error
	row := db.QueryRow(
		context.Background(),
		`INSERT INTO account (account_user_name, money_balance) VALUES ($1, $2) RETURNING (account_id, account_user_name, money_balance);`,
		account.Username, account.MoneyBalance)

	var createdAccount model.Account
	err = row.Scan(&createdAccount)

	return &createdAccount, err
}

func (accountRepository *AccountRepository) GetAccountById(accountId int) (*model.Account, error) {
	db := accountRepository.db

	var account model.Account
	err := pgxscan.Get(context.Background(), db, &account, `SELECT * FROM account WHERE account_id = $1;`, accountId)
	return &account, err
}

func (accountRepository *AccountRepository) UpdateAccount(account *model.Account) {
	accountRepository.db.QueryRow(
		context.Background(),
		"UPDATE account SET account_user_name = $1, money_balance = $2, connected_lobby_id = $3 WHERE account_id = $4;",
		account.Username, account.MoneyBalance, account.ConnectedLobbyId, account.ID)
}

func (accountRepository *AccountRepository) RemoveLobbyConnection(accountId int) {
	accountRepository.db.QueryRow(
		context.Background(),
		"UPDATE account SET connected_lobby_id = null WHERE account_id = $1;",
		accountId)
}
