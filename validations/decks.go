package validations

import (
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddDecks(c *gin.Context) (body models.Decks, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetDecksParams struct {
	IDs      string `form:"ids"`
	DeckOids []primitive.ObjectID
}
type GetDeckHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}

func GetDecks(c *gin.Context) (header GetDeckHeader, params GetDecksParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.DeckOids, err = ConvertToObjectIds(params.IDs)
	}
	return
}

type GetDeckUri struct {
	DeckId  string `uri:"deckId" binding:"required"`
	DeckOid primitive.ObjectID
}

func GetDeck(c *gin.Context) (uri GetDeckUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.DeckOid, err = primitive.ObjectIDFromHex(uri.DeckId)
	}
	return
}

type UpdateDeckUri struct {
	DeckId  string `uri:"deckId" binding:"required"`
	DeckOid primitive.ObjectID
}

func UpdateDeck(c *gin.Context) (uri UpdateDeckUri, body models.Deck, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.DeckOid, err = primitive.ObjectIDFromHex(uri.DeckId)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type DeleteDecksParams struct {
	IDs      *string `form:"ids" binding:"required"`
	DeckOids []primitive.ObjectID
}

func DeleteDecks(c *gin.Context) (params DeleteDecksParams, err error) {
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.DeckOids, err = ConvertToObjectIds(*params.IDs)
	}
	return
}

type DeleteDeckUri struct {
	DeckId  string `uri:"deckId" binding:"required"`
	DeckOid primitive.ObjectID
}

func DeleteDeck(c *gin.Context) (uri DeleteDeckUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.DeckOid, err = primitive.ObjectIDFromHex(uri.DeckId)
	}
	return
}
