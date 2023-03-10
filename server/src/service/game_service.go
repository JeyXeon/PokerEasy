package service

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/dto"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
	"strconv"
)

type LobbyChannels struct {
	ConnectedChannel    chan *dto.Event
	DisconnectedChannel chan *dto.Event
	ReadyStateChannel   chan *dto.Event
	GameEventsChannel   chan *dto.Event
}

func NewLobbyChannels() *LobbyChannels {
	lobbyChannels := new(LobbyChannels)
	lobbyChannels.ConnectedChannel = make(chan *dto.Event)
	lobbyChannels.DisconnectedChannel = make(chan *dto.Event)
	lobbyChannels.ReadyStateChannel = make(chan *dto.Event)
	return lobbyChannels
}

type GameService struct {
	LobbyEventPipesByLobbyIds   *common.RWLockerMap[*LobbyChannels]
	ConnectedAccountsByLobbyIds *common.RWLockerMap[[]*model.Account]
	ConnectedAccountIds         *common.RWLockerMap[bool]

	accountService common.AccountService
	lobbyService   common.LobbyService
}

func GetGameService(accountService common.AccountService, lobbyService common.LobbyService) *GameService {
	gameService := new(GameService)

	gameService.LobbyEventPipesByLobbyIds = common.NewRWLockerMap(make(map[int]*LobbyChannels))
	gameService.ConnectedAccountsByLobbyIds = common.NewRWLockerMap(make(map[int][]*model.Account))
	gameService.ConnectedAccountIds = common.NewRWLockerMap(make(map[int]bool))

	gameService.accountService = accountService
	gameService.lobbyService = lobbyService

	return gameService
}

func (gameService *GameService) ListenWebsocket(conn *websocket.Conn) {
	accountService := gameService.accountService

	lobbyEventPipesByLobbyIds := gameService.LobbyEventPipesByLobbyIds
	connectedAccountIds := gameService.ConnectedAccountIds

	lobbyIdParam, err := GetPathParamFromConnection(conn, "lobbyId")
	if err != nil {
		HandleError(conn, err, "LobbyIdParam param is not present")
		return
	}
	lobbyId, err := strconv.Atoi(lobbyIdParam)
	if err != nil {
		HandleError(conn, err, "Invalid lobbyId param")
		return
	}

	accountIdParam, err := GetQueryParamFromConnection(conn, "accountId")
	if err != nil {
		HandleError(conn, err, "AccountId param is not present")
		return
	}
	accountId, err := strconv.Atoi(accountIdParam)
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

	_, accountConnected := connectedAccountIds.Load(accountId)
	if !accountConnected {
		if err := gameService.addActivePlayer(lobbyId, account, conn); err != nil {
			HandleError(conn, err, fmt.Sprintf("Lobby with id %d doesn't exist", lobbyId))
			return
		}
	}

	for {
		_, messageContent, err := conn.ReadMessage()
		if err != nil {
			return
		}

		event := &dto.Event{
			Body:       string(messageContent),
			Connection: conn,
			Account:    account,
		}

		if event.Body == dto.StartGameEvent {
			lobbyEventPipes, _ := lobbyEventPipesByLobbyIds.Load(lobbyId)
			lobbyEventPipes.ReadyStateChannel <- event
		}
	}
}

func (gameService *GameService) disconnectPlayer(account *model.Account, lobbyId int, conn *websocket.Conn) {
	connectedAccountIds := gameService.ConnectedAccountIds
	lobbyEventPipesByLobbyIds := gameService.LobbyEventPipesByLobbyIds
	connectedAccountsByLobbyIds := gameService.ConnectedAccountsByLobbyIds

	connectedAccountIds.Delete(account.ID)

	connectedAccounts, _ := connectedAccountsByLobbyIds.Load(lobbyId)
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
		Account:    account,
	}
	lobbyEventPipes, _ := lobbyEventPipesByLobbyIds.Load(lobbyId)
	lobbyEventPipes.DisconnectedChannel <- playerDisconnectedEvent

	if err := conn.Close(); err != nil {
		logrus.Error(err)
		return
	}
}

func (gameService *GameService) addActivePlayer(lobbyId int, account *model.Account, conn *websocket.Conn) error {
	accountService := gameService.accountService

	lobbyEventPipesByLobbyIds := gameService.LobbyEventPipesByLobbyIds
	connectedAccountIds := gameService.ConnectedAccountIds
	connectedAccountsByLobbyIds := gameService.ConnectedAccountsByLobbyIds

	account.ConnectedLobbyId = &lobbyId
	accountService.UpdateAccount(account)
	connectedAccounts, _ := connectedAccountsByLobbyIds.Load(lobbyId)
	connectedAccountsByLobbyIds.Store(lobbyId, append(connectedAccounts, account))

	_, lobbyExists := lobbyEventPipesByLobbyIds.Load(lobbyId)
	if !lobbyExists {
		if err := gameService.startLobbyProcessing(lobbyId); err != nil {
			return err
		}
	}

	connectedAccountIds.Store(account.ID, true)

	playerConnectedEvent := &dto.Event{
		Body:       dto.PlayerConnectedEvent,
		Connection: conn,
		Account:    account,
	}
	lobbyEventPipes, _ := lobbyEventPipesByLobbyIds.Load(lobbyId)
	lobbyEventPipes.ConnectedChannel <- playerConnectedEvent

	return nil
}

func (gameService *GameService) startLobbyProcessing(lobbyId int) error {
	lobbyService := gameService.lobbyService
	lobbyEventPipesByLobbyIds := gameService.LobbyEventPipesByLobbyIds

	lobbyChannels := NewLobbyChannels()
	lobbyEventPipesByLobbyIds.Store(lobbyId, lobbyChannels)

	lobby, err := lobbyService.GetLobbyById(lobbyId)

	go gameService.ProcessLobby(lobby, lobbyChannels)

	return err
}

func (gameService *GameService) ProcessLobby(lobby *model.Lobby, lobbyChannels *LobbyChannels) {
	connectedChannel := lobbyChannels.ConnectedChannel
	disconnectedChannel := lobbyChannels.DisconnectedChannel
	readyStateChannel := lobbyChannels.ReadyStateChannel

	connections := make([]*websocket.Conn, 0)
	lobbyState := dto.NewLobbyState(lobby)
	lobbyId := lobby.ID

	for {
		select {
		case event := <-connectedChannel:
			gameService.processPlayerConnection(event, lobbyState, &connections)

		case event := <-disconnectedChannel:
			if len(connections) == 1 {
				gameService.LobbyEventPipesByLobbyIds.Delete(lobbyId)
				break
			} else {
				gameService.processPlayerDisconnection(event, lobbyState, &connections)
			}

		case event := <-readyStateChannel:
			allReady := gameService.processReadyStateChanging(event, lobbyState, &connections)
			if allReady {
				gameService.processGameStart(lobbyChannels, connections, lobbyState)
			}
		}
	}
}
