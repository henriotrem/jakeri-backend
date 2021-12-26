package routers

import (
	"jakeri-backend/controllers"
	"jakeri-backend/middleware"

	"github.com/gin-gonic/gin"
)

func profilesRoutes(profiles *gin.RouterGroup) {

	profiles.POST("", middleware.JWTAuthMiddleware(), controllers.AddProfiles)
	profiles.GET("", middleware.JWTAuthMiddleware(), controllers.GetProfiles)
	profiles.GET("/:profileId", middleware.JWTAuthMiddleware(), controllers.GetProfile)
	profiles.PUT("/:profileId", middleware.JWTAuthMiddleware(), controllers.UpdateProfile)
	profiles.DELETE("", middleware.JWTAuthMiddleware(), controllers.DeleteProfiles)
	profiles.DELETE("/:profileId", middleware.JWTAuthMiddleware(), controllers.DeleteProfile)

}
