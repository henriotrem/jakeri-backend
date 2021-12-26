package authorizations

import (
	"errors"
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProfiles(c *gin.Context, groupId *primitive.ObjectID) (tokenData TokenData, err error) {

	tokenData = GetAccessTokenData(c)

	group := &models.Group{}
	group.Get(groupId, nil)

	if !tokenData.GroupIds[*groupId] && !tokenData.Roles["admin"] && *group.Audit.CreatedBy != tokenData.UserID {
		err = errors.New("Forbidden")
	}
	return
}

func GetProfiles(c *gin.Context, groupId *primitive.ObjectID) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.GroupIds[*groupId] && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func GetProfile(c *gin.Context, groupId *primitive.ObjectID) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.GroupIds[*groupId] && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func UpdateProfile(c *gin.Context, groupId *primitive.ObjectID) (tokenData TokenData, err error) {
	tokenData = GetAccessTokenData(c)
	if !tokenData.GroupIds[*groupId] && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteProfiles(c *gin.Context, groupId *primitive.ObjectID) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.GroupIds[*groupId] && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}

func DeleteProfile(c *gin.Context, groupId *primitive.ObjectID) (err error) {
	tokenData := GetAccessTokenData(c)
	if !tokenData.GroupIds[*groupId] && !tokenData.Roles["admin"] {
		err = errors.New("Forbidden")
	}
	return
}
