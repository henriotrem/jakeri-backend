package routers

import (
	"jakeri-backend/controllers"

	"github.com/gin-gonic/gin"
)

func confirmationRoutes(confirmation *gin.RouterGroup) {

	confirmation.POST("", controllers.AddConfirmation)
	confirmation.POST("/new", controllers.NewConfirmation)
}
