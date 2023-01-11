package service

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/dto"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/gofiber/websocket/v2"
)

func (gameService *GameService) processPlayerConnection(event *dto.Event, connections *[]*websocket.Conn) {
	*connections = append(*connections, event.Connection)
	accountName := event.Connection.Query("accountName", "")
	for _, connection := range *connections {
		err := connection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Player %s connected to lobby", accountName)))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (gameService *GameService) processPlayerDisconnection(event *dto.Event, connections *[]*websocket.Conn) {
	accountName := event.Connection.Query("accountName", "")
	deleteIdx := -1
	for i, connection := range *connections {
		if connection == event.Connection {
			deleteIdx = i
		} else {
			err := connection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Player %s disconnected from lobby", accountName)))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	*connections = append((*connections)[:deleteIdx], (*connections)[deleteIdx+1:]...)
}

func (gameService *GameService) processGameStart(
	event *dto.Event,
	lobbyChannels *LobbyChannels,
	connections []*websocket.Conn,
	connectedAccounts []*model.Account,
) {
	gameEventsChannel := make(chan *dto.Event)
	lobbyChannels.GameEventsChannel = gameEventsChannel
	gameState := dto.NewGameState(connectedAccounts)

	go processGame(gameEventsChannel, gameState, connections)

	accountName := event.Connection.Query("accountName", "")
	for _, connection := range connections {
		err := connection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Game started by %s", accountName)))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func processGame(
	gameEventsChannel chan *dto.Event,
	gameState *dto.GameState,
	connections []*websocket.Conn,
) {
}
