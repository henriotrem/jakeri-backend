package authorizations

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func GetUser(c *gin.Context, userId string) (err error) {
	tokenData := GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func UpdateUser(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteUsers(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteUser(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func UpdatePassword(c *gin.Context, userId string) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if tokenData.UserID != userId && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}
