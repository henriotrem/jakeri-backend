package validations

import (
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddCardsUri struct {
	ProfileId string `uri:"profileId" binding:"required"`
}

func AddCards(c *gin.Context) (uri AddCardsUri, body models.Cards, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetCardsHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetCardsUri struct {
	ProfileId string `uri:"profileId" binding:"required"`
}
type GetCardsParams struct {
	IDs      string `form:"ids"`
	CardOids []primitive.ObjectID
}

func GetCards(c *gin.Context) (header GetCardsHeader, uri GetCardsUri, params GetCardsParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.CardOids, err = ConvertToObjectIds(params.IDs)
	}
	return
}

type GetCardHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetCardUri struct {
	ProfileId string `uri:"profileId" binding:"required"`
	CardId    string `uri:"cardId" binding:"required"`
	CardOid   primitive.ObjectID
}

func GetCard(c *gin.Context) (header GetCardHeader, uri GetCardUri, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.CardOid, err = primitive.ObjectIDFromHex(uri.CardId)
	}
	return
}

type UpdateCardUri struct {
	ProfileId string `uri:"profileId" binding:"required"`
	CardId    string `uri:"cardId" binding:"required"`
	CardOid   primitive.ObjectID
}

func UpdateCard(c *gin.Context) (uri UpdateCardUri, body models.Card, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.CardOid, err = primitive.ObjectIDFromHex(uri.CardId)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type DeleteCardsUri struct {
	ProfileId string `uri:"profileId" binding:"required"`
}
type DeleteCardsParams struct {
	IDs      *string `form:"ids" binding:"required"`
	CardOids []primitive.ObjectID
}

func DeleteCards(c *gin.Context) (uri DeleteCardsUri, params DeleteCardsParams, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.CardOids, err = ConvertToObjectIds(*params.IDs)
	}
	return
}

type DeleteCardUri struct {
	ProfileId string `uri:"profileId" binding:"required"`
	CardId    string `uri:"cardId" binding:"required"`
	CardOid   primitive.ObjectID
}

func DeleteCard(c *gin.Context) (uri DeleteCardUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.CardOid, err = primitive.ObjectIDFromHex(uri.CardId)
	}
	return
}
