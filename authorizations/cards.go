package authorizations

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func AddCards(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId {
		err = errors.New("Forbidden")
	}
	return
}

func GetCards(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId {
		err = errors.New("Forbidden")
	}
	return
}

func UpdateCard(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteCards(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteCard(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId {
		err = errors.New("Forbidden")
	}
	return
}
