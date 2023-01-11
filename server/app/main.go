package main

import (
	"github.com/JeyXeon/poker-easy/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	accountHandlers := handlers.GetAccountHandlers()
	lobbyHandlers := handlers.GetLobbyHandlers()

	webApp := fiber.New()

	account := webApp.Group("/account")
	account.Post("", accountHandlers.CreateAccountHandler)
	account.Get("/:accountId", accountHandlers.GetAccountHandler)

	lobby := webApp.Group("/lobby")
	lobby.Post("", lobbyHandlers.CreateLobby)
	lobby.Get("/all", lobbyHandlers.GetAllLobbies)
	lobby.Get("/:lobbyId", lobbyHandlers.GetLobbyById)
	lobby.Get("/:lobbyId/connect", websocket.New(lobbyHandlers.ConnectToLobby))

	logrus.Fatal(webApp.Listen(":80"))
}
