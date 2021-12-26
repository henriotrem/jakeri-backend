package routers

import (
	"jakeri-backend/controllers"

	"github.com/gin-gonic/gin"
)

func sessionRoutes(session *gin.RouterGroup) {

	session.POST("", controllers.AddSession)
	session.PUT("", controllers.UpdateSession)
}
