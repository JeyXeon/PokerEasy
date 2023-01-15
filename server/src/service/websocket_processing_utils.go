package service

import (
	"errors"
	"fmt"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"
)

func GetQueryParamFromConnection(connection *websocket.Conn, paramName string) (string, error) {
	param := connection.Query(paramName, "")
	if param == "" {
		return param, errors.New("query param not found")
	}
	return param, nil
}

func GetPathParamFromConnection(connection *websocket.Conn, paramName string) (string, error) {
	param := connection.Params(paramName, "")
	if param == "" {
		return param, errors.New("path param not found")
	}
	return param, nil
}

func HandleError(conn *websocket.Conn, err error, message string) {
	logrus.WithError(err).Info(err.Error())
	if err := conn.Close(); err != nil {
		logrus.Error(err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(message))); err != nil {
		logrus.Error(err)
	}
}
