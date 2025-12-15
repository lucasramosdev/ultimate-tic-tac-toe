package manager

import (
	"sync"

	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/client"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/rooms"
)

type GameManager struct {
	Rooms      map[string]*rooms.Room
	Register   chan *client.Client
	Unregister chan *client.Client
	mu         sync.Mutex
}

type InitMsg struct {
	Type   string `json:"type"`
	Player string `json:"player"`
}
