package routers

import (
	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.Engine) {
	api := router.Group("")
	{
		usersRoutes(api.Group("/users"))
		confirmationRoutes(api.Group("/confirmation"))
		sessionRoutes(api.Group("/session"))
		recoveryRoutes(api.Group("/recovery"))
		postsRoutes(api.Group("/posts"))
		reviewsRoutes(api.Group("/reviews"))
		groupsRoutes(api.Group("/groups"))
		statusRoutes(api.Group("/status"))
	}
}