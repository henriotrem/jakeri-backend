package routers

import (
	"jakeri-backend/controllers"
	"jakeri-backend/middleware"

	"github.com/gin-gonic/gin"
)

func usersRoutes(users *gin.RouterGroup) {

	users.POST("", controllers.AddUsers)
	users.GET("", middleware.JWTAuthMiddleware(), controllers.GetUsers)
	users.GET("/:userId", middleware.JWTAuthMiddleware(), controllers.GetUser)
	users.PUT("/:userId", middleware.JWTAuthMiddleware(), controllers.UpdateUser)
	users.DELETE("", middleware.JWTAuthMiddleware(), controllers.DeleteUsers)
	users.DELETE("/:userId", middleware.JWTAuthMiddleware(), controllers.DeleteUser)
	users.PUT("/:userId/password", middleware.JWTAuthMiddleware(), controllers.UpdatePassword)

}
