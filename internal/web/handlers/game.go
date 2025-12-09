package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/hash"
)

func PlayHanldler(c *gin.Context) {
	gameToken := hash.CreateGameToken()
	c.Redirect(http.StatusTemporaryRedirect, "/game/"+*gameToken)
}
