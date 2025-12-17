package rooms

import (
	"encoding/json"

	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/client"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/state"
)

func (r *Room) Run() {
	for message := range r.Broadcast {
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

func (r *Room) UnregisterClient(c *client.Client) {
	if r.PlayerX == c {
		r.PlayerX = nil
		close(c.Send)
	}

	if r.PlayerO == c {
		r.PlayerO = nil
		close(c.Send)
	}

	r.GameState = state.NewGameState()
	r.BroadcastState()
}

func (r *Room) HandleMessage() {
	for msg := range r.Incoming {
		switch msg.Type {
		case "move":
			if r.GameState.MakeMove(msg.OuterIndex, msg.InnerIndex, msg.Player) {
				r.BroadcastState()
			}
		case "reset":
			r.GameState = state.NewGameState()
			r.BroadcastState()
		}
	}
}

func (r *Room) BroadcastState() {
	var playerXName, playerOName string
	if r.PlayerX != nil {
		playerXName = r.PlayerX.Name
	}
	if r.PlayerO != nil {
		playerOName = r.PlayerO.Name
	}

	stateJSON, _ := json.Marshal(struct {
		Type        string           `json:"type"`
		State       *state.GameState `json:"state"`
		PlayerXName string           `json:"playerXName"`
		PlayerOName string           `json:"playerOName"`
	}{
		Type:        "state",
		State:       r.GameState,
		PlayerXName: playerXName,
		PlayerOName: playerOName,
	})
	r.Broadcast <- stateJSON
}
