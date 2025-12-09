package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
