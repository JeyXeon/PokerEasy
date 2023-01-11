package service

import (
	"github.com/JeyXeon/poker-easy/model"
	"github.com/JeyXeon/poker-easy/repository"
)

type AccountRepository interface {
	CreateAccount(account model.Account) (*model.Account, error)
	GetAccountById(accountId int) *model.Account
	UpdateAccount(account *model.Account)
	RemoveLobbyConnection(accountId int)
}

type AccountService struct {
	accountRepository AccountRepository
}

func GetAccountService() *AccountService {
	accountRepository := repository.GetAccountRepository()
	return &AccountService{accountRepository: accountRepository}
}

func (accountService *AccountService) SaveNewAccount(accountDto model.Account) *model.Account {
	accountRepository := accountService.accountRepository
	createdAccount, err := accountRepository.CreateAccount(accountDto)
	if err != nil {

	}
	return createdAccount
}

func (accountService *AccountService) GetAccountById(userId int) *model.Account {
	accountRepository := accountService.accountRepository
	existingAccount := accountRepository.GetAccountById(userId)
	return existingAccount
}

func (accountService *AccountService) UpdateAccount(account *model.Account) {
	accountService.accountRepository.UpdateAccount(account)
}

func (accountService *AccountService) RemoveLobbyConnection(accountId int) {
	accountService.accountRepository.RemoveLobbyConnection(accountId)
}
