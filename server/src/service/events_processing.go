package service

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/dto"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
)

func (gameService *GameService) processPlayerConnection(event *dto.Event, lobbyState *dto.LobbyState, connections *[]*websocket.Conn) {
	*connections = append(*connections, event.Connection)
	lobbyState.AddPlayer(dto.AccountToPlayer(event.Account))

	accountName := event.Account.Username
	for _, connection := range *connections {
		responseMessage := dto.NewLobbyEventResponse(fmt.Sprintf("Player %s connected to lobby", accountName), *lobbyState)
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
			responseMessage := dto.NewLobbyEventResponse(fmt.Sprintf("Player %s disconnected from lobby", account.Username), *lobbyState)
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
	lobbyState *dto.LobbyState,
) {
	gameEventsChannel := make(chan *dto.Event)
	lobbyChannels.GameEventsChannel = gameEventsChannel
	gameState := dto.NewGameState()
	lobbyState.GameState = gameState

	go gameService.processGame(gameEventsChannel, lobbyState, connections)
}

func (gameService *GameService) processReadyStateChanging(event *dto.Event, lobbyState *dto.LobbyState, connections *[]*websocket.Conn) bool {
	account := event.Account
	allReady := lobbyState.ChangeReadyState(account.ID)
	for _, connection := range *connections {
		responseMessage := dto.NewLobbyEventResponse(fmt.Sprintf("Player %s is ready", account.Username), *lobbyState)
		err := connection.WriteMessage(websocket.TextMessage, responseMessage.ToJson())
		if err != nil {
			logrus.Error(err)
		}
	}
	return allReady
}

func (gameService *GameService) processGame(
	gameEventsChannel chan *dto.Event,
	lobbyState *dto.LobbyState,
	connections []*websocket.Conn,
) {
	participantPlaces := make([]int, 0)
	handsByPlaces := make(map[int][]dto.PlayingCard)
	for place, player := range lobbyState.ConnectedPlayersByPlaces {
		if player != nil && player.IsReady {
			player.IsGameMember = true
			participantPlaces = append(participantPlaces, place)
			handsByPlaces[place] = make([]dto.PlayingCard, 0, 2)
		}
	}
	sort.Slice(participantPlaces, func(i, j int) bool {
		return participantPlaces[i] < participantPlaces[j]
	})

	gameState := lobbyState.GameState
	cardCounter := 0

	for rowCounter := 0; rowCounter < 2; rowCounter++ {
		for _, place := range participantPlaces {
			handsByPlaces[place] = append(handsByPlaces[place], gameState.Deck[cardCounter])
			cardCounter++
		}
	}

	for _, connection := range connections {
		accountIdParam, _ := GetQueryParamFromConnection(connection, "accountId")
		accountId, _ := strconv.Atoi(accountIdParam)

		personalLobbyState := *lobbyState
		playerPlace := personalLobbyState.PlacesByConnectedPlayerIds[accountId]
		hand, exist := handsByPlaces[playerPlace]
		if exist && personalLobbyState.ConnectedPlayersByPlaces[playerPlace].ID == accountId {
			personalLobbyState.ConnectedPlayersByPlaces[playerPlace].Hand = &hand
		}

		responseMessage := dto.NewLobbyEventResponse("Game started", personalLobbyState)
		err := connection.WriteMessage(websocket.TextMessage, responseMessage.ToJson())
		if err != nil {
			logrus.Error(err)
		}

		if exist && personalLobbyState.ConnectedPlayersByPlaces[playerPlace].ID == accountId {
			personalLobbyState.ConnectedPlayersByPlaces[playerPlace].Hand = nil
		}
	}
}
