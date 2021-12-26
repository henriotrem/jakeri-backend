package routers

import (
	"jakeri-backend/controllers"
	"jakeri-backend/middleware"

	"github.com/gin-gonic/gin"
)

func groupsRoutes(groups *gin.RouterGroup) {

	groups.POST("", middleware.JWTAuthMiddleware(), controllers.AddGroups)
	groups.GET("", middleware.JWTAuthMiddleware(), controllers.GetGroups)
	groups.GET("/:groupId", middleware.JWTAuthMiddleware(), controllers.GetGroup)
	groups.PUT("/:groupId", middleware.JWTAuthMiddleware(), controllers.UpdateGroup)
	groups.DELETE("", middleware.JWTAuthMiddleware(), controllers.DeleteGroups)
	groups.DELETE("/:groupId", middleware.JWTAuthMiddleware(), controllers.DeleteGroup)

	profilesRoutes(groups.Group("/:groupId/profiles"))
}
