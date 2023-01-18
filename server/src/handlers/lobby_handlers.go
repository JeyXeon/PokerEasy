package handlers

import (
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"net/http"
	"strconv"
)

type LobbyHandlers struct {
	lobbyService common.LobbyService
	gameService  common.GameService
}

func GetLobbyHandlers(lobbyService common.LobbyService, gameService common.GameService) *LobbyHandlers {
	lobbyHandlers := new(LobbyHandlers)
	lobbyHandlers.lobbyService = lobbyService
	lobbyHandlers.gameService = gameService
	return lobbyHandlers
}

type CreateLobbyRequest struct {
	LobbyName  string `json:"lobbyName"`
	MaxPlayers int    `json:"maxPlayers"`
	CreatorId  int    `json:"creatorId"`
}

func (lobbyHandlers *LobbyHandlers) WsUpgradeCheck(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (lobbyHandlers *LobbyHandlers) CreateLobby(c *fiber.Ctx) error {
	var request CreateLobbyRequest
	if err := c.BodyParser(&request); err != nil {
		return SendResponse(c, nil, err, fiber.StatusBadRequest)
	}
	creatorId := request.CreatorId
	lobby := model.Lobby{LobbyName: request.LobbyName, MaxPlayers: request.MaxPlayers, CreatorId: creatorId}

	lobbyService := lobbyHandlers.lobbyService
	createdLobby, err := lobbyService.SaveNewLobby(lobby)

	return SendResponse(c, createdLobby, err, fiber.StatusBadRequest)
}

func (lobbyHandlers *LobbyHandlers) GetLobbyById(c *fiber.Ctx) error {
	lobbyIdParam := c.Params("lobbyId", "")
	if lobbyIdParam == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	lobbyId, err := strconv.Atoi(lobbyIdParam)
	if err != nil || lobbyId < 0 {
		return SendResponse(c, nil, err, fiber.StatusUnprocessableEntity)
	}

	lobbyService := lobbyHandlers.lobbyService
	existingLobby, err := lobbyService.GetLobbyById(lobbyId)

	return SendResponse(c, existingLobby, err, fiber.StatusNotFound)
}

func (lobbyHandlers *LobbyHandlers) GetAllLobbies(c *fiber.Ctx) error {
	lobbyService := lobbyHandlers.lobbyService
	allLobbies, err := lobbyService.GetAllLobbies()

	return SendResponse(c, allLobbies, err, fiber.StatusInternalServerError)
}

func (lobbyHandlers *LobbyHandlers) ConnectToLobby(c *websocket.Conn) {
	gameService := lobbyHandlers.gameService
	gameService.ListenWebsocket(c)
}
