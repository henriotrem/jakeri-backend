package routers

import (
	"jakeri-backend/controllers"
	"jakeri-backend/middleware"

	"github.com/gin-gonic/gin"
)

func postsRoutes(posts *gin.RouterGroup) {

	posts.POST("", middleware.JWTAuthMiddleware(), controllers.AddPosts)
	posts.GET("", middleware.JWTAuthMiddleware(), controllers.GetPosts)
	posts.GET("/:postId", middleware.JWTAuthMiddleware(), controllers.GetPost)
	posts.PUT("/:postId", middleware.JWTAuthMiddleware(), controllers.UpdatePost)
	posts.DELETE("", middleware.JWTAuthMiddleware(), controllers.DeletePosts)
	posts.DELETE("/:postId", middleware.JWTAuthMiddleware(), controllers.DeletePost)
}
