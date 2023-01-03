package main

import (
	"github.com/JeyXeon/poker-easy/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()
	webApp.Post("/account", handlers.AddUserHandler)
	webApp.Get("/account/:accountId", handlers.GetUserHandler)

	logrus.Fatal(webApp.Listen(":80"))
}
