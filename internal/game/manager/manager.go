package manager

import (
	"encoding/json"
	"log"

	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/client"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/rooms"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/state"
)

var Instance = &GameManager{
	Rooms:      make(map[string]*rooms.Room),
	Register:   make(chan *client.Client),
	Unregister: make(chan *client.Client),
}

func (m *GameManager) Run() {
	for {
		select {
		case c := <-m.Register:
			m.mu.Lock()
			room, exists := m.Rooms[c.Code]
			if !exists {
				room = &rooms.Room{
					Code:      c.Code,
					GameState: state.NewGameState(),
					Broadcast: make(chan []byte),
					Incoming:  make(chan client.Message),
				}
				m.Rooms[c.Code] = room
				go room.Run()
				go room.HandleMessage()
			}

			if room.PlayerX != nil && room.PlayerO != nil {
				log.Println("Room full")
				close(c.Send)
				m.mu.Unlock()
				continue
			}

			if room.PlayerX == nil && room.PlayerO != nil {
				room.PlayerX = c
				c.Player = "X"
			}

			if room.PlayerO == nil {
				room.PlayerO = c
				c.Player = "O"
			}

			c.Room = room.Incoming

			initMsg, _ := json.Marshal(InitMsg{
				Type:   "init",
				Player: c.Player,
			})

			c.Send <- initMsg

			room.BroadcastState()
			m.mu.Unlock()

		case client := <-m.Unregister:
			m.mu.Lock()
			if room, ok := m.Rooms[client.Code]; ok {
				room.UnregisterClient(client)
			}
			m.mu.Unlock()
		}
	}
}

func (m *GameManager) GetRoom(code string) *rooms.Room {
	m.mu.Lock()
	defer m.mu.Unlock()
	if room, ok := m.Rooms[code]; ok {
		return room
	}
	room := &rooms.Room{
		Code:      code,
		GameState: state.NewGameState(),
		Broadcast: make(chan []byte),
		Incoming:  make(chan client.Message),
	}
	m.Rooms[code] = room
	go room.Run()
	go room.HandleMessage()
	return room
}
