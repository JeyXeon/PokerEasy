package service

import (
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/model"
)

type AccountService struct {
	accountRepository common.AccountRepository
}

func GetAccountService(accountRepository common.AccountRepository) *AccountService {
	accountService := new(AccountService)
	accountService.accountRepository = accountRepository
	return accountService
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
