package routers

import (
	"jakeri-backend/controllers"

	"github.com/gin-gonic/gin"
)

func statusRoutes(status *gin.RouterGroup) {

	status.GET("/health", controllers.GetHealth)
}
