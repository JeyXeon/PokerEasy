package handlers

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/JeyXeon/poker-easy/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
		return fmt.Errorf("body parser: %w", err)
	}

	accountService := accountHandlers.accountService
	accountData := model.Account{Username: request.UserName, MoneyBalance: request.MoneyBalance}
	createdAccount := accountService.SaveNewAccount(accountData)
	return c.JSON(createdAccount)
}

func (accountHandlers *AccountHandlers) GetAccountHandler(c *fiber.Ctx) error {
	accountIdParam := c.Params("accountId", "")
	if accountIdParam == "" {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	accountId, err := strconv.Atoi(accountIdParam)
	if err != nil || accountId < 0 {
		return fmt.Errorf("invalid account id: %w", err)
	}

	accountService := accountHandlers.accountService
	existingAccount := accountService.GetAccountById(accountId)
	return c.JSON(existingAccount)
}
