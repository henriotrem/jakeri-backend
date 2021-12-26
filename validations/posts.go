package validations

import (
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPosts(c *gin.Context) (body models.Posts, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetPostsParams struct {
	IDs      string `form:"ids"`
	PostOids []primitive.ObjectID
}
type GetPostHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}

func GetPosts(c *gin.Context) (header GetPostHeader, params GetPostsParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.PostOids, err = ConvertToObjectIds(params.IDs)
	}
	return
}

type GetPostUri struct {
	PostId  string `uri:"postId" binding:"required"`
	PostOid primitive.ObjectID
}

func GetPost(c *gin.Context) (uri GetPostUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.PostOid, err = primitive.ObjectIDFromHex(uri.PostId)
	}
	return
}

type UpdatePostUri struct {
	PostId  string `uri:"postId" binding:"required"`
	PostOid primitive.ObjectID
}

func UpdatePost(c *gin.Context) (uri UpdatePostUri, body models.Post, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.PostOid, err = primitive.ObjectIDFromHex(uri.PostId)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type DeletePostsParams struct {
	IDs      *string `form:"ids" binding:"required"`
	PostOids []primitive.ObjectID
}

func DeletePosts(c *gin.Context) (params DeletePostsParams, err error) {
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.PostOids, err = ConvertToObjectIds(*params.IDs)
	}
	return
}

type DeletePostUri struct {
	PostId  string `uri:"postId" binding:"required"`
	PostOid primitive.ObjectID
}

func DeletePost(c *gin.Context) (uri DeletePostUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.PostOid, err = primitive.ObjectIDFromHex(uri.PostId)
	}
	return
}
