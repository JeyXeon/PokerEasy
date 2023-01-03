package handlers

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/JeyXeon/poker-easy/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type AccountService interface {
	SaveNewUser(accountDto model.Account) model.Account
	GetUserById(userId int) model.Account
}

type AccountHandlers struct {
	accountService AccountService
}

func GetAccountHandlers() *AccountHandlers {
	accountService := service.GetAccountService()
	return &AccountHandlers{accountService: accountService}
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
	createdAccount := accountService.SaveNewUser(accountData)

	return c.SendString(fmt.Sprintf("Created account: %s.", createdAccount.ToString()))
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
	existingAccount := accountService.GetUserById(accountId)

	return c.SendString(fmt.Sprintf("Found account: %s.", existingAccount.ToString()))
}
