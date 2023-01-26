package dto

import "encoding/json"

type ResponseMessage struct {
	Message string      `json:"message"`
	State   interface{} `json:"state"`
}

func NewLobbyEventResponse(message string, lobbyState *LobbyState) *ResponseMessage {
	responseMessage := new(ResponseMessage)
	responseMessage.Message = message
	responseMessage.State = lobbyState

	return responseMessage
}

func NewGameEventResponse(message string, gameState *GameState) *ResponseMessage {
	responseMessage := new(ResponseMessage)
	responseMessage.Message = message
	responseMessage.State = gameState

	return responseMessage
}

func (responseMessage *ResponseMessage) ToJson() []byte {
	responseJson, _ := json.Marshal(responseMessage)
	return responseJson
}
