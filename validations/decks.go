package validations

import (
	"jakeri-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddDecksUri struct {
	UserId  string `uri:"userId" binding:"required"`
	UserOid primitive.ObjectID
}

func AddDecks(c *gin.Context) (uri AddDecksUri, body models.Decks, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.UserOid, err = primitive.ObjectIDFromHex(uri.UserId)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetDecksHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetDecksUri struct {
	UserId  string `uri:"userId" binding:"required"`
	UserOid primitive.ObjectID
}
type GetDecksParams struct {
	IDs      string `form:"ids"`
	DeckOids []primitive.ObjectID
}

func GetDecks(c *gin.Context) (header GetDecksHeader, uri GetDecksUri, params GetDecksParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.UserOid, err = primitive.ObjectIDFromHex(uri.UserId)
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
	UserId  string `uri:"userId" binding:"required"`
	UserOid primitive.ObjectID
	DeckId  string `uri:"deckId" binding:"required"`
	DeckOid primitive.ObjectID
}

func GetDeck(c *gin.Context) (uri GetDeckUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.UserOid, err = primitive.ObjectIDFromHex(uri.UserId)
	}
	if err == nil {
		uri.DeckOid, err = primitive.ObjectIDFromHex(uri.DeckId)
	}
	return
}

type UpdateDeckUri struct {
	UserId  string `uri:"userId" binding:"required"`
	UserOid primitive.ObjectID
	DeckId  string `uri:"deckId" binding:"required"`
	DeckOid primitive.ObjectID
}

func UpdateDeck(c *gin.Context) (uri UpdateDeckUri, body models.Deck, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.UserOid, err = primitive.ObjectIDFromHex(uri.UserId)
	}
	if err == nil {
		uri.DeckOid, err = primitive.ObjectIDFromHex(uri.DeckId)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type DeleteDecksUri struct {
	UserId  string `uri:"userId" binding:"required"`
	UserOid primitive.ObjectID
}
type DeleteDecksParams struct {
	IDs      *string `form:"ids" binding:"required"`
	DeckOids []primitive.ObjectID
}

func DeleteDecks(c *gin.Context) (uri UpdateDeckUri, params DeleteDecksParams, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.UserOid, err = primitive.ObjectIDFromHex(uri.UserId)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil {
		params.DeckOids, err = ConvertToObjectIds(*params.IDs)
	}
	return
}

type DeleteDeckUri struct {
	UserId  string `uri:"userId" binding:"required"`
	UserOid primitive.ObjectID
	DeckId  string `uri:"deckId" binding:"required"`
	DeckOid primitive.ObjectID
}

func DeleteDeck(c *gin.Context) (uri DeleteDeckUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		uri.UserOid, err = primitive.ObjectIDFromHex(uri.UserId)
	}
	if err == nil {
		uri.DeckOid, err = primitive.ObjectIDFromHex(uri.DeckId)
	}
	return
}
