package authorizations

import (
	"github.com/gin-gonic/gin"
)

func AddDecks(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func GetDecks(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func GetDeck(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func UpdateDeck(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func DeleteDecks(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func DeleteDeck(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}
