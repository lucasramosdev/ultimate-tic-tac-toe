package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/game/manager"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/web/routes"
)

func main() {
	go manager.Instance.Run()
	r := gin.Default()

	r.Static("/css", "web/css")
	r.LoadHTMLGlob("web/templates/*")
	routes.Setup(r)
	r.Run()
}
