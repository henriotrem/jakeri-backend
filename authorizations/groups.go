package authorizations

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func AddReviews(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func GetReviews(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func GetReview(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}
