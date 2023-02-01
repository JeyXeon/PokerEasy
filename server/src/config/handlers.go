package config

import "github.com/JeyXeon/poker-easy/handlers"

type AppHandlers struct {
	AccountHandlers *handlers.AccountHandlers
	LobbyHandlers   *handlers.LobbyHandlers
}

func GetAppHandlers(appServices *AppServices) *AppHandlers {
	appHandlers := new(AppHandlers)
	appHandlers.AccountHandlers = handlers.GetAccountHandlers(appServices.accountService)
	appHandlers.LobbyHandlers = handlers.GetLobbyHandlers(appServices.lobbyService, appServices.gameService)

	return appHandlers
}
