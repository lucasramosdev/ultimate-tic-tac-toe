package rooms

import (
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/client"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/state"
)

type Room struct {
	Code      string
	PlayerX   *client.Client
	PlayerO   *client.Client
	GameState *state.GameState
	Broadcast chan []byte
	Incoming  chan client.Message
}
