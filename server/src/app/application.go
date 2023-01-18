package main

import (
	"github.com/JeyXeon/poker-easy/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Application struct {
	fiberApp *fiber.App

	db *pgxpool.Pool

	AppServices     *config.AppServices
	AppHandlers     *config.AppHandlers
	AppRepositories *config.AppRepositories
}

func NewApplication() *Application {
	db := config.GetDbConnection()
	appRepositories := config.GetAppRepositories(db)
	appServices := config.GetAppServices(appRepositories)
	appHandlers := config.GetAppHandlers(appServices)

	application := new(Application)
	application.fiberApp = fiber.New()
	application.db = db
	application.AppRepositories = appRepositories
	application.AppHandlers = appHandlers
	application.AppServices = appServices

	return application
}

func (application *Application) configureRoutes() {
	webApp := application.fiberApp

	appHandlers := application.AppHandlers
	accountHandlers := appHandlers.AccountHandlers
	lobbyHandlers := appHandlers.LobbyHandlers

	account := webApp.Group("/account")
	account.Post("", accountHandlers.CreateAccountHandler)
	account.Get("/:accountId", accountHandlers.GetAccountHandler)

	lobby := webApp.Group("/lobby")
	lobby.Post("", lobbyHandlers.CreateLobby)
	lobby.Get("/all", lobbyHandlers.GetAllLobbies)
	lobby.Get("/:lobbyId", lobbyHandlers.GetLobbyById)
	lobby.Get("/:lobbyId/connect", websocket.New(lobbyHandlers.ConnectToLobby))
}

func (application *Application) startWebApp() {
	webApp := application.fiberApp
	webApp.Use(recover.New())
	logrus.Fatal(webApp.Listen(":8000"))
}
