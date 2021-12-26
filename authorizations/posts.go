package authorizations

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func AddPosts(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func UpdatePost(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeletePosts(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeletePost(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}
