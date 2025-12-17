package client

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Room   chan Message
	Code   string
	Conn   *websocket.Conn
	Send   chan []byte
	Player string
	Name   string
}

type Message struct {
	Type       string `json:"type"`
	Payload    string `json:"payload"`
	OuterIndex int    `json:"outerIndex,omitempty"`
	InnerIndex int    `json:"innerIndex,omitempty"`
	Room       string `json:"room"`
	Player     string `json:"player,omitempty"`
}
