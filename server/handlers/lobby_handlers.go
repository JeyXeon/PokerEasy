package handlers

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/common"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/JeyXeon/poker-easy/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"net/http"
	"strconv"
)

type LobbyHandlers struct {
	lobbyService common.LobbyService
	gameService  common.GameService
}

func GetLobbyHandlers() *LobbyHandlers {
	lobbyHandlers := new(LobbyHandlers)
	lobbyHandlers.lobbyService = service.GetLobbyService()
	lobbyHandlers.gameService = service.GetGameService()
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
		return fmt.Errorf("body parser: %w", err)
	}
	creatorId := request.CreatorId
	lobby := model.Lobby{LobbyName: request.LobbyName, MaxPlayers: request.MaxPlayers, CreatorId: creatorId}

	lobbyService := lobbyHandlers.lobbyService
	createdLobby := lobbyService.SaveNewLobby(lobby)
	return c.JSON(createdLobby)
}

func (lobbyHandlers *LobbyHandlers) GetLobbyById(c *fiber.Ctx) error {
	lobbyIdParam := c.Params("lobbyId", "")
	if lobbyIdParam == "" {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	lobbyId, err := strconv.Atoi(lobbyIdParam)
	if err != nil || lobbyId < 0 {
		return fmt.Errorf("invalid account id: %w", err)
	}

	lobbyService := lobbyHandlers.lobbyService
	existingLobby := lobbyService.GetLobbyById(lobbyId)

	return c.JSON(existingLobby)
}

func (lobbyHandlers *LobbyHandlers) GetAllLobbies(c *fiber.Ctx) error {
	lobbyService := lobbyHandlers.lobbyService
	allLobbies := lobbyService.GetAllLobbies()

	return c.JSON(allLobbies)
}

func (lobbyHandlers *LobbyHandlers) ConnectToLobby(c *websocket.Conn) {
	gameService := lobbyHandlers.gameService
	gameService.ListenWebsocket(c)
}
