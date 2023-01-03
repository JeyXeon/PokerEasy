package repository

import (
	"github.com/JeyXeon/poker-easy/config"
	"github.com/JeyXeon/poker-easy/model"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func GetAccountRepository() *AccountRepository {
	dbCon := config.GetDbConnection()
	return &AccountRepository{db: dbCon}
}

func (accountRepository *AccountRepository) CreateAccount(account model.Account) (model.Account, error) {
	db := accountRepository.db
	tx := db.Begin()
	if err := tx.Create(&account).Error; err != nil {
		tx.Rollback()
		return model.Account{}, err
	}
	tx.Select("account_id", &account).Where("name = ?", account.Username)
	tx.Commit()

	return account, nil
}

func (accountRepository *AccountRepository) GetAccountById(accountId int) model.Account {
	account := model.Account{}
	db := accountRepository.db
	db.First(&account, "account_id = ?", accountId)
	return account
}
