package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SendResponse(c *fiber.Ctx, body interface{}, err error, errStatus int) error {
	if err != nil {
		logrus.WithError(err).Info(err.Error())
		return c.Status(errStatus).SendString(err.Error())
	}
	return c.JSON(body)
}
