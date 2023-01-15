package handlers

import (
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/JeyXeon/poker-easy/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type AccountHandlers struct {
	accountService common.AccountService
}

func GetAccountHandlers() *AccountHandlers {
	accountHandlers := new(AccountHandlers)
	accountHandlers.accountService = service.GetAccountService()
	return accountHandlers
}

type CreateAccountRequest struct {
	UserName     string `json:"userName"`
	MoneyBalance int64  `json:"moneyBalance"`
}

func (accountHandlers *AccountHandlers) CreateAccountHandler(c *fiber.Ctx) error {
	var request CreateAccountRequest
	if err := c.BodyParser(&request); err != nil {
		return SendResponse(c, nil, err, fiber.StatusBadRequest)
	}

	accountService := accountHandlers.accountService
	accountData := model.Account{Username: request.UserName, MoneyBalance: request.MoneyBalance}
	createdAccount, err := accountService.SaveNewAccount(accountData)

	return SendResponse(c, createdAccount, err, fiber.StatusBadRequest)
}

func (accountHandlers *AccountHandlers) GetAccountHandler(c *fiber.Ctx) error {
	accountIdParam := c.Params("accountId", "")
	if accountIdParam == "" {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	accountId, err := strconv.Atoi(accountIdParam)
	if err != nil || accountId < 0 {
		return SendResponse(c, nil, err, fiber.StatusUnprocessableEntity)
	}

	accountService := accountHandlers.accountService
	existingAccount, err := accountService.GetAccountById(accountId)

	return SendResponse(c, existingAccount, err, fiber.StatusNotFound)
}
