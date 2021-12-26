package validations

import (
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddGroups(c *gin.Context) (body models.Groups, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetGroupsHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetGroupsParams struct {
	IDs       string `form:"ids"`
	GroupOids []primitive.ObjectID
}

func GetGroups(c *gin.Context) (header GetGroupsHeader, params GetGroupsParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.GroupOids, err = ConvertToObjectIds(params.IDs)
	}
	return
}

type GetGroupHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetGroupUri struct {
	GroupId  string `uri:"groupId" binding:"required"`
	GroupOid primitive.ObjectID
}

func GetGroup(c *gin.Context) (header GetGroupHeader, uri GetGroupUri, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.GroupOid, err = primitive.ObjectIDFromHex(uri.GroupId)
	}
	return
}

type UpdateGroupUri struct {
	GroupId  string `uri:"groupId" binding:"required"`
	GroupOid primitive.ObjectID
}

func UpdateGroup(c *gin.Context) (uri UpdateGroupUri, body models.Group, err error) {
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

type DeleteGroupsParams struct {
	IDs       *string `form:"ids" binding:"required"`
	GroupOids []primitive.ObjectID
}

func DeleteGroups(c *gin.Context) (params DeleteGroupsParams, err error) {
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.GroupOids, err = ConvertToObjectIds(*params.IDs)
	}
	return
}

type DeleteGroupUri struct {
	GroupId  string `uri:"groupId" binding:"required"`
	GroupOid primitive.ObjectID
}

func DeleteGroup(c *gin.Context) (uri DeleteGroupUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.GroupOid, err = primitive.ObjectIDFromHex(uri.GroupId)
	}
	return
}
