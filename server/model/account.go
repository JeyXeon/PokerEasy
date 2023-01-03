package model

import "fmt"

type Account struct {
	ID             int    `gorm:"column:account_id"`
	Username       string `gorm:"column:account_user_name"`
	ConnectedLobby int    `gorm:"foreignKey:UserRefer;references:LobbyId"`
	MoneyBalance   int64
}

func (Account) TableName() string {
	return "account"
}

func (account Account) ToString() string {
	return fmt.Sprintf("{ID: %d, Username: %s, MoneyBalance: %d}", account.ID, account.Username, account.MoneyBalance)
}
