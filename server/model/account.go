package model

type Account struct {
	ID               int    `json:"accountId" db:"account_id"`
	Username         string `json:"username" db:"account_user_name"`
	MoneyBalance     int64  `json:"moneyBalance" db:"money_balance"`
	ConnectedLobbyId int    `json:"connectedLobbyIds" db:"connected_lobby_id"`
}
