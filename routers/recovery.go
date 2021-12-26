package routers

import (
	"jakeri-backend/controllers"

	"github.com/gin-gonic/gin"
)

func recoveryRoutes(recovery *gin.RouterGroup) {

	recovery.POST("", controllers.AddRecovery)
	recovery.POST("/new", controllers.NewRecovery)
}
