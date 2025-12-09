package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/web/handlers"
)

func PingRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)
}
