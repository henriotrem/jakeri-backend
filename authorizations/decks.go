package authorizations

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func AddDecks(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func GetDecks(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func GetDeck(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func UpdateDeck(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteDecks(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteDeck(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}
