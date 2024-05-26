package service

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/dto"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
)

func (gameService *GameService) processPlayerConnection(event *dto.Event, lobbyId int) {
	connections, _ := gameService.ConnectionsByLobbyIds.Load(lobbyId)
	*connections = append(*connections, event.Connection)

	lobbyState, _ := gameService.StatesByLobbyIds.Load(lobbyId)
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

func (gameService *GameService) processPlayerDisconnection(event *dto.Event, lobbyId int) bool {
	account := event.Account
	lobbyState, _ := gameService.StatesByLobbyIds.Load(lobbyId)
	lobbyState.RemovePlayer(account.ID)
	connections, _ := gameService.ConnectionsByLobbyIds.Load(lobbyId)
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

	if len(*connections) == 0 {
		gameService.LobbyEventPipesByLobbyIds.Delete(lobbyId)
		return true
	}

	return false
}

func (gameService *GameService) processReadyStateChanging(event *dto.Event, lobbyId int) {
	account := event.Account
	lobbyState, _ := gameService.StatesByLobbyIds.Load(lobbyId)
	lobbyState.ChangeReadyState(account.ID)
	connections, _ := gameService.ConnectionsByLobbyIds.Load(lobbyId)
	for _, connection := range *connections {
		responseMessage := dto.NewLobbyEventResponse(fmt.Sprintf("Player %s is ready", account.Username), *lobbyState)
		err := connection.WriteMessage(websocket.TextMessage, responseMessage.ToJson())
		if err != nil {
			logrus.Error(err)
		}
	}

	if len(*connections) > 1 && lobbyState.AllPlayersAreReady() {
		gameService.processGameStart(lobbyId)
	}
}

func (gameService *GameService) processGameStart(lobbyId int) {
	lobbyState, _ := gameService.StatesByLobbyIds.Load(lobbyId)

	gameState := dto.NewGameState()
	lobbyState.GameState = gameState

	participantPlaces := make([]int, 0)
	for place, player := range lobbyState.ConnectedPlayersByPlaces {
		if player != nil && player.IsReady {
			player.IsGameMember = true
			participantPlaces = append(participantPlaces, place)
			playerHand := make([]dto.PlayingCard, 2, 2)
			player.Hand = &playerHand
			player.IsReady = false
		}
	}
	sort.Slice(participantPlaces, func(i, j int) bool {
		return participantPlaces[i] < participantPlaces[j]
	})

	for _, place := range participantPlaces {
		for rowCounter := 0; rowCounter < 2; rowCounter++ {
			playerHand := *lobbyState.ConnectedPlayersByPlaces[place].Hand
			playerHand[rowCounter] = gameState.Deck[gameState.CardCounter]
			gameState.CardCounter++
		}
	}

	connections, _ := gameService.ConnectionsByLobbyIds.Load(lobbyId)
	for _, connection := range *connections {
		gameService.sendPersonalLobbyStateWithMessage(connection, lobbyState, "GAME STARTED")
	}
}

func (gameService *GameService) processCheckEvent(event *dto.Event, lobbyId int) {
	lobbyState, _ := gameService.StatesByLobbyIds.Load(lobbyId)
	if lobbyState.AllPlayersAreReady() {
		gameService.processNextGameRound(lobbyState)
	}

	connections, _ := gameService.ConnectionsByLobbyIds.Load(lobbyId)
	for _, connection := range *connections {
		gameService.sendPersonalLobbyStateWithMessage(connection, lobbyState, "NEXT ROUND")
	}
}

func (gameService *GameService) processNextGameRound(state *dto.LobbyState) {
	for _, player := range state.ConnectedPlayersByPlaces {
		if player != nil {
			player.IsReady = false
		}
	}

	gameState := state.GameState
	var counter int
	if gameState.CurrentRound == 0 {
		counter = 3
	} else {
		counter = 1
	}
	for i := 0; i < counter; i++ {
		gameState.CardsOnTable = append(gameState.CardsOnTable, gameState.Deck[gameState.CardCounter])
		gameState.CardCounter++
	}
}

func (gameService *GameService) sendPersonalLobbyStateWithMessage(
	connection *websocket.Conn,
	lobbyState *dto.LobbyState,
	message string,
) {
	accountIdParam, _ := GetQueryParamFromConnection(connection, "accountId")
	accountId, _ := strconv.Atoi(accountIdParam)

	personalLobbyState := *lobbyState
	personalLobbyState.ConnectedPlayersByPlaces = make(map[int]*dto.Player, len(lobbyState.ConnectedPlayersByPlaces))

	for place, player := range lobbyState.ConnectedPlayersByPlaces {
		if player != nil && player.ID != accountId {
			playerCopy := *player
			playerCopy.Hand = nil
			personalLobbyState.ConnectedPlayersByPlaces[place] = &playerCopy
		} else {
			personalLobbyState.ConnectedPlayersByPlaces[place] = player
		}
	}

	responseMessage := dto.NewLobbyEventResponse(message, personalLobbyState)
	err := connection.WriteMessage(websocket.TextMessage, responseMessage.ToJson())
	if err != nil {
		logrus.Error(err)
	}
}
