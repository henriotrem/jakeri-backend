package validations

import (
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddCards(c *gin.Context) (body models.Cards, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetCardsHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetCardsParams struct {
	IDs      string `form:"ids"`
	CardOids []primitive.ObjectID
}

func GetCards(c *gin.Context) (header GetCardsHeader, params GetCardsParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
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
	CardId  string `uri:"cardId" binding:"required"`
	CardOid primitive.ObjectID
}

func GetCard(c *gin.Context) (uri GetCardUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.CardOid, err = primitive.ObjectIDFromHex(uri.CardId)
	}
	return
}

type UpdateCardUri struct {
	CardId  string `uri:"cardId" binding:"required"`
	CardOid primitive.ObjectID
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

type DeleteCardsParams struct {
	IDs      *string `form:"ids" binding:"required"`
	CardOids []primitive.ObjectID
}

func DeleteCards(c *gin.Context) (params DeleteCardsParams, err error) {
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.CardOids, err = ConvertToObjectIds(*params.IDs)
	}
	return
}

type DeleteCardUri struct {
	CardId  string `uri:"cardId" binding:"required"`
	CardOid primitive.ObjectID
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
