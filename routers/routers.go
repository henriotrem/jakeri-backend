package routers

import (
	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.Engine) {
	api := router.Group("")
	{
		usersRoutes(api.Group("/users"))
		profilesRoutes(api.Group("/profiles"))
		confirmationRoutes(api.Group("/confirmation"))
		sessionRoutes(api.Group("/session"))
		recoveryRoutes(api.Group("/recovery"))
		statusRoutes(api.Group("/status"))
	}
}
