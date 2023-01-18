package config

import "github.com/JeyXeon/poker-easy/service"

type AppServices struct {
	accountService *service.AccountService
	lobbyService   *service.LobbyService
	gameService    *service.GameService
}

func GetAppServices(appRepositories *AppRepositories) *AppServices {
	appServices := new(AppServices)
	appServices.accountService = service.GetAccountService(appRepositories.accountRepository)
	appServices.lobbyService = service.GetLobbyService(appRepositories.lobbyRepository)
	appServices.gameService = service.GetGameService(appServices.accountService, appServices.lobbyService)

	return appServices
}
