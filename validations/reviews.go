package validations

import (
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddReviews(c *gin.Context) (body models.Reviews, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetReviewsParams struct {
	IDs         string `form:"ids"`
	ReviewsOids []primitive.ObjectID
}
type GetReviewsHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}

func GetReviews(c *gin.Context) (header GetReviewsHeader, params GetReviewsParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.ReviewsOids, err = ConvertToObjectIds(params.IDs)
	}
	return
}

type GetReviewUri struct {
	ReviewId  string `uri:"reviewId" binding:"required"`
	ReviewOid primitive.ObjectID
}
type GetReviewHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}

func GetReview(c *gin.Context) (header GetReviewHeader, uri GetReviewUri, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.ReviewOid, err = primitive.ObjectIDFromHex(uri.ReviewId)
	}
	return
}
