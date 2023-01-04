package handlers

import (
	"fmt"
	"github.com/JeyXeon/poker-easy/model"
	"github.com/JeyXeon/poker-easy/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"net/http"
	"strconv"
)

type LobbyService interface {
	SaveNewLobby(lobby model.Lobby) model.Lobby
	GetLobbyById(lobbyId int) model.Lobby
	GetAllLobbies() model.Lobbies
}

type LobbyHandlers struct {
	lobbyService LobbyService
}

func GetLobbyHandlers() *LobbyHandlers {
	lobbyService := service.GetLobbyService()
	return &LobbyHandlers{lobbyService: lobbyService}
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
	return c.SendString(fmt.Sprintf("Created lobby: %s.", createdLobby.ToString()))
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

	return c.SendString(fmt.Sprintf("Found lobby: %s.", existingLobby.ToString()))
}

func (lobbyHandlers *LobbyHandlers) GetAllLobbies(c *fiber.Ctx) error {
	lobbyService := lobbyHandlers.lobbyService
	allLobbies := lobbyService.GetAllLobbies()

	return c.SendString(fmt.Sprintf("Found lobbies: %s.", allLobbies.ToString()))
}
