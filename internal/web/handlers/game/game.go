package game

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/client"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/hash"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/manager"
)

func PlayHandler(c *gin.Context) {
	var playRequest PlayRequest
	if err := c.ShouldBind(&playRequest); err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	if playRequest.RoomCode == "" {
		gameToken := *hash.CreateGameToken()
		c.Redirect(http.StatusFound, "/game/"+gameToken)
		return
	}

	if !hash.ValidateGameToken(playRequest.RoomCode) {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.Redirect(http.StatusFound, "/game/"+playRequest.RoomCode)

}

func ServeGame(c *gin.Context) {
	roomCode := c.Param("token")
	if !hash.ValidateGameToken(roomCode) {
		c.Redirect(http.StatusFound, "/")
		return
	}

	manager.Instance.GetRoom(roomCode)

	c.HTML(http.StatusOK, "game.tmpl", gin.H{
		"RoomCode": roomCode,
	})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(c *gin.Context) {
	roomCode := c.Param("token")
	if !hash.ValidateGameToken(roomCode) {
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade ws: %v", err)
		return
	}

	client := &client.Client{Conn: conn, Send: make(chan []byte, 256), Code: roomCode}
	manager.Instance.GetRoom(roomCode)
	manager.Instance.Register <- client

	go client.WritePump()
	go func() {
		client.ReadPump()
		manager.Instance.Unregister <- client
	}()

}
