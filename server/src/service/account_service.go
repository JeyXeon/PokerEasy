package service

import (
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/JeyXeon/poker-easy/repository"
)

type AccountService struct {
	accountRepository common.AccountRepository
}

func GetAccountService() *AccountService {
	accountRepository := repository.GetAccountRepository()
	return &AccountService{accountRepository: accountRepository}
}

func (accountService *AccountService) SaveNewAccount(accountDto model.Account) (*model.Account, error) {
	accountRepository := accountService.accountRepository
	createdAccount, err := accountRepository.CreateAccount(accountDto)
	return createdAccount, err
}

func (accountService *AccountService) GetAccountById(userId int) (*model.Account, error) {
	accountRepository := accountService.accountRepository
	existingAccount, err := accountRepository.GetAccountById(userId)
	return existingAccount, err
}

func (accountService *AccountService) UpdateAccount(account *model.Account) {
	accountService.accountRepository.UpdateAccount(account)
}

func (accountService *AccountService) RemoveLobbyConnection(accountId int) {
	accountService.accountRepository.RemoveLobbyConnection(accountId)
}
