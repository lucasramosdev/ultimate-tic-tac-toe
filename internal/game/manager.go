package game

import (
	"encoding/json"
	"log"
	"sync"
)

type GameManager struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	mu         sync.Mutex
}

var Instance = &GameManager{
	Rooms:      make(map[string]*Room),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

type Room struct {
	Code      string
	PlayerX   *Client
	PlayerO   *Client
	GameState *GameState
	broadcast chan []byte
}

func (m *GameManager) Run() {
	for {
		select {
		case client := <-m.Register:
			m.mu.Lock()
			room, exists := m.Rooms[client.Room.Code]
			if !exists {
				room = &Room{
					Code:      client.Room.Code,
					GameState: NewGameState(),
					broadcast: make(chan []byte),
				}
				m.Rooms[client.Room.Code] = room
				go room.Run()
			}

			if room.PlayerX == nil {
				room.PlayerX = client
				client.Player = "X"
			} else if room.PlayerO == nil {
				room.PlayerO = client
				client.Player = "O"
			} else {
				log.Println("Room full")
				close(client.Send)
				m.mu.Unlock()
				continue
			}

			initMsg, _ := json.Marshal(struct {
				Type   string `json:"type"`
				Player string `json:"player"`
			}{
				Type:   "init",
				Player: client.Player,
			})
			client.Send <- initMsg

			room.BroadcastState()
			m.mu.Unlock()

		case client := <-m.Unregister:
			if client.Room != nil {
				client.Room.UnregisterClient(client)
			}
		}
	}
}

func (r *Room) Run() {
	for message := range r.broadcast {
		if r.PlayerX != nil {
			select {
			case r.PlayerX.Send <- message:
			default:
				close(r.PlayerX.Send)
				r.PlayerX = nil
			}
		}
		if r.PlayerO != nil {
			select {
			case r.PlayerO.Send <- message:
			default:
				close(r.PlayerO.Send)
				r.PlayerO = nil
			}
		}
	}
}

func (r *Room) UnregisterClient(c *Client) {
	if r.PlayerX == c {
		r.PlayerX = nil
		close(c.Send)
	} else if r.PlayerO == c {
		r.PlayerO = nil
		close(c.Send)
	}

	r.GameState = NewGameState()
	r.BroadcastState()
}

func (r *Room) HandleMessage(c *Client, msg Message) {
	switch msg.Type {
	case "move":
		if r.GameState.MakeMove(msg.OuterIndex, msg.InnerIndex, c.Player) {
			r.BroadcastState()
		}
	case "reset":
		r.GameState = NewGameState()
		r.BroadcastState()
	}
}

func (r *Room) BroadcastState() {
	stateJSON, _ := json.Marshal(struct {
		Type  string     `json:"type"`
		State *GameState `json:"state"`
	}{
		Type:  "state",
		State: r.GameState,
	})
	r.broadcast <- stateJSON
}

func (m *GameManager) GetRoom(code string) *Room {
	m.mu.Lock()
	defer m.mu.Unlock()
	if room, ok := m.Rooms[code]; ok {
		return room
	}
	room := &Room{
		Code:      code,
		GameState: NewGameState(),
		broadcast: make(chan []byte),
	}
	m.Rooms[code] = room
	go room.Run()
	return room
}
