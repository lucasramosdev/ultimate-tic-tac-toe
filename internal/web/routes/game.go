package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasramosdev/jogo-da-velha-dois/internal/web/handlers"
)

func GameRoutes(r *gin.Engine) {
	r.POST("/play", handlers.PlayHanldler)
}
