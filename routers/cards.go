package routers

import (
	"jakeri-backend/controllers"
	"jakeri-backend/middleware"

	"github.com/gin-gonic/gin"
)

func cardsRoutes(cards *gin.RouterGroup) {

	cards.POST("", middleware.JWTAuthMiddleware(), controllers.AddCards)
	cards.GET("", middleware.JWTAuthMiddleware(), controllers.GetCards)
	cards.GET("/:cardId", middleware.JWTAuthMiddleware(), controllers.GetCard)
	cards.PUT("/:cardId", middleware.JWTAuthMiddleware(), controllers.UpdateCard)
	cards.DELETE("", middleware.JWTAuthMiddleware(), controllers.DeleteCards)
	cards.DELETE("/:cardId", middleware.JWTAuthMiddleware(), controllers.DeleteCard)
}
