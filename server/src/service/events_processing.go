package service

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/dto"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
)

func (gameService *GameService) processPlayerConnection(event *dto.Event, lobbyState *dto.LobbyState, connections *[]*websocket.Conn) {
	*connections = append(*connections, event.Connection)
	lobbyState.AddPlayer(dto.AccountToPlayer(event.Account))

	accountName := event.Account.Username
	for _, connection := range *connections {
		responseMessage := dto.NewLobbyEventResponse(fmt.Sprintf("Player %s connected to lobby", accountName), lobbyState)
		err := connection.WriteMessage(websocket.TextMessage, responseMessage.ToJson())
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (gameService *GameService) processPlayerDisconnection(event *dto.Event, lobbyState *dto.LobbyState, connections *[]*websocket.Conn) {
	account := event.Account
	lobbyState.RemovePlayer(account.ID)
	var deleteIdx int
	for i, connection := range *connections {
		if connection == event.Connection {
			deleteIdx = i
		} else {
			responseMessage := dto.NewLobbyEventResponse(fmt.Sprintf("Player %s disconnected from lobby", account.Username), lobbyState)
			err := connection.WriteMessage(websocket.TextMessage, responseMessage.ToJson())
			if err != nil {
				logrus.Error(err)
			}
		}
	}
	*connections = append((*connections)[:deleteIdx], (*connections)[deleteIdx+1:]...)
}

func (gameService *GameService) processGameStart(
	lobbyChannels *LobbyChannels,
	connections []*websocket.Conn,
	playersByPlaces map[int]*dto.Player,
) {
	gameEventsChannel := make(chan *dto.Event)
	lobbyChannels.GameEventsChannel = gameEventsChannel
	gameState := dto.NewGameState(playersByPlaces)

	go processGame(gameEventsChannel, gameState, connections)

	for _, connection := range connections {
		err := connection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Game started")))
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (gameService *GameService) processReadyStateChanging(event *dto.Event, lobbyState *dto.LobbyState, connections *[]*websocket.Conn) bool {
	account := event.Account
	allReady := lobbyState.ChangeReadyState(account.ID)
	for _, connection := range *connections {
		responseMessage := dto.NewLobbyEventResponse(fmt.Sprintf("Player %s is ready", account.Username), lobbyState)
		err := connection.WriteMessage(websocket.TextMessage, responseMessage.ToJson())
		if err != nil {
			logrus.Error(err)
		}
	}
	return allReady
}

func processGame(
	gameEventsChannel chan *dto.Event,
	gameState *dto.GameState,
	connections []*websocket.Conn,
) {
}
