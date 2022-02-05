package routers

import (
	"jakeri-backend/controllers"
	"jakeri-backend/middleware"

	"github.com/gin-gonic/gin"
)

func decksRoutes(decks *gin.RouterGroup) {

	decks.POST("", middleware.JWTAuthMiddleware(), controllers.AddDecks)
	decks.GET("", middleware.JWTAuthMiddleware(), controllers.GetDecks)
	decks.GET("/:deckId", middleware.JWTAuthMiddleware(), controllers.GetDeck)
	decks.PUT("/:deckId", middleware.JWTAuthMiddleware(), controllers.UpdateDeck)
	decks.DELETE("", middleware.JWTAuthMiddleware(), controllers.DeleteDecks)
	decks.DELETE("/:deckId", middleware.JWTAuthMiddleware(), controllers.DeleteDeck)
}
