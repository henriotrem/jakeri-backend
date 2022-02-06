package controllers

import (
	"jakeri-backend/authorizations"
	"jakeri-backend/models"
	"jakeri-backend/utils"
	"jakeri-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddDecks(c *gin.Context) {

	uri, decks, err := validations.AddDecks(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.AddDecks(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	ids, err := decks.Add(&tokenData.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(ids) == 0 {
		c.Status(http.StatusConflict)
	} else if len(ids) < len(decks) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(ids))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(ids))
	}
}

func GetDecks(c *gin.Context) {

	header, uri, params, err := validations.GetDecks(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.GetDecks(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	decks := models.Decks{}
	err = decks.Get(params.DeckOids, &tokenData.UserID, header.Embed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(decks) == 0 && len(params.DeckOids) > 0 {
		c.Status(http.StatusNotFound)
	} else if len(decks) < len(params.DeckOids) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(decks))
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(decks))
	}
}

func GetDeck(c *gin.Context) {

	header, uri, err := validations.GetDeck(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.GetDeck(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	deck := &models.Deck{}
	err = deck.Get(&uri.DeckOid, &tokenData.UserID, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if deck == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(deck))
	}
}

func UpdateDeck(c *gin.Context) {

	uri, deck, err := validations.UpdateDeck(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.UpdateDeck(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	modified, matched, err := deck.Update(&uri.DeckOid, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if modified == 0 && matched == 0 {
		c.Status(http.StatusNotFound)
	} else if modified == 0 && matched == 1 {
		c.Status(http.StatusNotModified)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(deck))
	}
}

func DeleteDecks(c *gin.Context) {
	uri, params, err := validations.DeleteDecks(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.DeleteDecks(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	decks := models.Decks{}
	res, err := decks.Delete(params.DeckOids, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else if res < len(params.DeckOids) {
		c.Status(http.StatusMultiStatus)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func DeleteDeck(c *gin.Context) {

	uri, err := validations.DeleteDeck(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.DeleteDeck(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	deck := &models.Deck{}

	res, err := deck.Delete(&uri.DeckOid, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusNoContent)
	}
}
