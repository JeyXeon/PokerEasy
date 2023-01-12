package service

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/dto"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/gofiber/websocket/v2"
	"log"
	"strconv"
)

type LobbyChannels struct {
	ConnectedChannel    chan *dto.Event
	DisconnectedChannel chan *dto.Event
	GameStartChannel    chan *dto.Event
	GameEventsChannel   chan *dto.Event
}

func NewLobbyChannels() *LobbyChannels {
	lobbyChannels := new(LobbyChannels)
	lobbyChannels.ConnectedChannel = make(chan *dto.Event)
	lobbyChannels.DisconnectedChannel = make(chan *dto.Event)
	lobbyChannels.GameStartChannel = make(chan *dto.Event)
	return lobbyChannels
}

type GameService struct {
	LobbyEventPipesByLobbyIds   map[int]*LobbyChannels
	ConnectedAccountsByLobbyIds map[int][]*model.Account
	ConnectedAccountIds         map[int]bool

	accountService common.AccountService
	lobbyService   common.LobbyService
}

func GetGameService() *GameService {
	gameService := new(GameService)

	gameService.LobbyEventPipesByLobbyIds = make(map[int]*LobbyChannels)
	gameService.ConnectedAccountIds = make(map[int]bool)
	gameService.ConnectedAccountsByLobbyIds = make(map[int][]*model.Account)

	gameService.accountService = GetAccountService()
	gameService.lobbyService = GetLobbyService()

	return gameService
}

func (gameService *GameService) ListenWebsocket(conn *websocket.Conn) {
	accountService := gameService.accountService

	lobbyEventPipesByLobbyIds := gameService.LobbyEventPipesByLobbyIds
	connectedAccountIds := gameService.ConnectedAccountIds

	lobbyId, err := strconv.Atoi(conn.Params("lobbyId", ""))
	if err != nil {
		HandleError(conn, err, "Invalid lobbyId param")
		return
	}

	accountId, err := strconv.Atoi(conn.Query("accountId", ""))
	if err != nil {
		HandleError(conn, err, "Invalid accountId param")
		return
	}
	account, err := accountService.GetAccountById(accountId)
	if err != nil {
		HandleError(conn, err, fmt.Sprintf("Account with id %d doesn't exist", accountId))
		return
	}

	defer gameService.disconnectPlayer(account, lobbyId, conn)

	_, accountConnected := connectedAccountIds[accountId]
	if !accountConnected {
		gameService.addActivePlayer(lobbyId, account, conn)
	}

	for {
		_, messageContent, err := conn.ReadMessage()
		if err != nil {
			return
		}

		event := &dto.Event{
			Body:       string(messageContent),
			Connection: conn,
		}

		if event.Body == dto.StartGameEvent {
			lobbyEventPipesByLobbyIds[lobbyId].GameStartChannel <- event
		}
	}
}

func (gameService *GameService) disconnectPlayer(account *model.Account, lobbyId int, conn *websocket.Conn) {
	connectedAccountIds := gameService.ConnectedAccountIds
	lobbyEventPipesByLobbyIds := gameService.LobbyEventPipesByLobbyIds
	connectedAccountsByLobbyIds := gameService.ConnectedAccountsByLobbyIds

	delete(connectedAccountIds, account.ID)

	connectedAccounts := connectedAccountsByLobbyIds[lobbyId]
	var deleteIdx int
	for i, connectedAccount := range connectedAccounts {
		if connectedAccount.ID == account.ID {
			deleteIdx = i
		}
	}
	connectedAccounts = append(connectedAccounts[:deleteIdx], connectedAccounts[deleteIdx+1:]...)

	gameService.accountService.RemoveLobbyConnection(account.ID)

	playerDisconnectedEvent := &dto.Event{
		Body:       dto.PlayerDisconnectedEvent,
		Connection: conn,
	}
	lobbyEventPipesByLobbyIds[lobbyId].DisconnectedChannel <- playerDisconnectedEvent

	if err := conn.Close(); err != nil {
		log.Println(err)
		return
	}
}

func (gameService *GameService) addActivePlayer(lobbyId int, account *model.Account, conn *websocket.Conn) {
	accountService := gameService.accountService
	lobbyEventPipesByLobbyIds := gameService.LobbyEventPipesByLobbyIds
	connectedAccountIds := gameService.ConnectedAccountIds
	connectedAccountsByLobbyIds := gameService.ConnectedAccountsByLobbyIds

	account.ConnectedLobbyId = lobbyId
	accountService.UpdateAccount(account)
	connectedAccountsByLobbyIds[lobbyId] = append(connectedAccountsByLobbyIds[lobbyId], account)

	_, lobbyExists := lobbyEventPipesByLobbyIds[lobbyId]
	if !lobbyExists {
		gameService.startLobbyProcessing(lobbyId)
	}

	connectedAccountIds[account.ID] = true

	playerConnectedEvent := &dto.Event{
		Body:       dto.PlayerConnectedEvent,
		Connection: conn,
	}
	lobbyEventPipesByLobbyIds[lobbyId].ConnectedChannel <- playerConnectedEvent
}

func (gameService *GameService) startLobbyProcessing(lobbyId int) {

	lobbyEventPipesByLobbyIds := gameService.LobbyEventPipesByLobbyIds

	lobbyChannels := NewLobbyChannels()
	lobbyEventPipesByLobbyIds[lobbyId] = lobbyChannels

	go gameService.ProcessLobby(lobbyId, lobbyChannels)
}

func (gameService *GameService) ProcessLobby(lobbyId int, lobbyChannels *LobbyChannels) {
	connectedChannel := lobbyChannels.ConnectedChannel
	disconnectedChannel := lobbyChannels.DisconnectedChannel
	gameStartChannel := lobbyChannels.GameStartChannel

	connections := make([]*websocket.Conn, 0)

	for {
		select {
		case event := <-connectedChannel:
			gameService.processPlayerConnection(event, &connections)

		case event := <-disconnectedChannel:
			if len(connections) == 1 {
				delete(gameService.LobbyEventPipesByLobbyIds, lobbyId)
				break
			} else {
				gameService.processPlayerDisconnection(event, &connections)
			}

		case event := <-gameStartChannel:
			accounts := gameService.ConnectedAccountsByLobbyIds[lobbyId]
			gameService.processGameStart(event, lobbyChannels, connections, accounts)
		}
	}
}
