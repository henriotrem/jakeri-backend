package validations

import (
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddProfilesUri struct {
	GroupId  string `uri:"groupId" binding:"required"`
	GroupOid primitive.ObjectID
}

func AddProfiles(c *gin.Context) (uri AddProfilesUri, body models.Profiles, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.GroupOid, err = primitive.ObjectIDFromHex(uri.GroupId)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetProfilesHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetProfilesUri struct {
	GroupId  string `uri:"groupId" binding:"required"`
	GroupOid primitive.ObjectID
}
type GetProfilesParams struct {
	IDs         string `form:"ids"`
	ProfileOids []primitive.ObjectID
}

func GetProfiles(c *gin.Context) (header GetProfilesHeader, uri GetProfilesUri, params GetProfilesParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.GroupOid, err = primitive.ObjectIDFromHex(uri.GroupId)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.ProfileOids, err = ConvertToObjectIds(params.IDs)
	}
	return
}

type GetProfileHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetProfileUri struct {
	GroupId    string `uri:"groupId" binding:"required"`
	GroupOid   primitive.ObjectID
	ProfileId  string `uri:"profileId" binding:"required"`
	ProfileOid primitive.ObjectID
}

func GetProfile(c *gin.Context) (header GetProfileHeader, uri GetProfileUri, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.GroupOid, err = primitive.ObjectIDFromHex(uri.GroupId)
	}
	if err == nil {
		uri.ProfileOid, err = primitive.ObjectIDFromHex(uri.ProfileId)
	}
	return
}

type UpdateProfileUri struct {
	GroupId    string `uri:"groupId" binding:"required"`
	GroupOid   primitive.ObjectID
	ProfileId  string `uri:"profileId" binding:"required"`
	ProfileOid primitive.ObjectID
}

func UpdateProfile(c *gin.Context) (uri UpdateProfileUri, body models.Profile, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.GroupOid, err = primitive.ObjectIDFromHex(uri.GroupId)
	}
	if err == nil {
		uri.ProfileOid, err = primitive.ObjectIDFromHex(uri.ProfileId)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type DeleteProfilesUri struct {
	GroupId  string `uri:"groupId" binding:"required"`
	GroupOid primitive.ObjectID
}
type DeleteProfilesParams struct {
	IDs         string `form:"ids" binding:"required"`
	ProfileOids []primitive.ObjectID
}

func DeleteProfiles(c *gin.Context) (uri DeleteProfilesUri, params DeleteProfilesParams, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.GroupOid, err = primitive.ObjectIDFromHex(uri.GroupId)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.ProfileOids, err = ConvertToObjectIds(params.IDs)
	}
	return
}

type DeleteProfileUri struct {
	GroupId    string `uri:"groupId" binding:"required"`
	GroupOid   primitive.ObjectID
	ProfileId  string `uri:"profileId" binding:"required"`
	ProfileOid primitive.ObjectID
}

func DeleteProfile(c *gin.Context) (uri DeleteProfileUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.GroupOid, err = primitive.ObjectIDFromHex(uri.GroupId)
	}
	if err == nil {
		uri.ProfileOid, err = primitive.ObjectIDFromHex(uri.ProfileId)
	}
	return
}
