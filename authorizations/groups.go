package authorizations

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddGroups(c *gin.Context) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func GetGroups(c *gin.Context, groupOids *[]primitive.ObjectID) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func GetGroup(c *gin.Context, groupId *primitive.ObjectID) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func UpdateGroup(c *gin.Context, groupId *primitive.ObjectID) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteGroups(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteGroup(c *gin.Context) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}
