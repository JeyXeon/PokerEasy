package dto

import (
	"github.com/JeyXeon/poker-easy/model"
	"github.com/gofiber/websocket/v2"
)

const (
	StartGameEvent          = "START_GAME"
	PlayerConnectedEvent    = "PLAYER_CONNECTED"
	PlayerDisconnectedEvent = "PLAYER_DISCONNECTED"
	PlayerCheckEvent        = "PLAYER_CHECK"
)

type Event struct {
	Body       string
	Connection *websocket.Conn
	Account    *model.Account
}
