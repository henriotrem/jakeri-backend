package routers

import (
	"jakeri-backend/controllers"
	"jakeri-backend/middleware"

	"github.com/gin-gonic/gin"
)

func reviewsRoutes(reviews *gin.RouterGroup) {

	reviews.POST("", middleware.JWTAuthMiddleware(), controllers.AddReviews)
	reviews.GET("", middleware.JWTAuthMiddleware(), controllers.GetReviews)
	reviews.GET("/:reviewId", middleware.JWTAuthMiddleware(), controllers.GetReview)
}
