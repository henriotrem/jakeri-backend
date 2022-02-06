package routers

import (
	"jakeri-backend/controllers"
	"jakeri-backend/middleware"

	"github.com/gin-gonic/gin"
)

func profilesRoutes(profiles *gin.RouterGroup) {

	profiles.GET("", middleware.JWTAuthMiddleware(), controllers.GetProfiles)
	profiles.GET("/:profileId", middleware.JWTAuthMiddleware(), controllers.GetProfile)

	cardsRoutes(profiles.Group("/:profileId/cards"))
}
