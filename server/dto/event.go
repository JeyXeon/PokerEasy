package dto

import "github.com/gofiber/websocket/v2"

const (
	StartGameEvent          = "START_GAME"
	PlayerConnectedEvent    = "PLAYER_CONNECTED"
	PlayerDisconnectedEvent = "PLAYER_DISCONNECTED"
)

type Event struct {
	Body       string
	Connection *websocket.Conn
}
