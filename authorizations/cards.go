package authorizations

import (
	"github.com/gin-gonic/gin"
)

func AddCards(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func UpdateCard(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func DeleteCards(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}

func DeleteCard(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	return
}
