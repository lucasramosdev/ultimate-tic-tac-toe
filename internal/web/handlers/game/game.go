package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/hash"
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

	game.Instance.GetRoom(roomCode)

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
		return
	}

	client := &game.Client{Manager: game.Instance, Conn: conn, Send: make(chan []byte, 256)}
	client.Room = game.Instance.GetRoom(roomCode)

	client.Manager.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
