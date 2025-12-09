package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/hash"
)

func main() {
	r := gin.Default()

	r.Static("/css", "web/css")
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{})
	})

	r.POST("/new-game", func(c *gin.Context) {
		gameToken := hash.CreateGameToken()
		c.JSON(http.StatusCreated, gin.H{
			"token": *gameToken,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
