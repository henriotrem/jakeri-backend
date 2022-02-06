package controllers

import (
	"jakeri-backend/authorizations"
	"jakeri-backend/models"
	"jakeri-backend/utils"
	"jakeri-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCards(c *gin.Context) {

	uri, cards, err := validations.AddCards(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.AddCards(c, uri.ProfileId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	ids, err := cards.Add(&tokenData.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(ids) == 0 {
		c.Status(http.StatusConflict)
	} else if len(ids) < len(cards) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(ids))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(ids))
	}
}

func GetCards(c *gin.Context) {

	header, uri, params, err := validations.GetCards(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	cards := models.Cards{}
	err = cards.Get(params.CardOids, &uri.ProfileId, header.Embed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(cards) == 0 && len(params.CardOids) > 0 {
		c.Status(http.StatusNotFound)
	} else if len(cards) < len(params.CardOids) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(cards))
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(cards))
	}
}

func GetCard(c *gin.Context) {

	header, uri, err := validations.GetCard(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	card := &models.Card{}
	err = card.Get(&uri.CardOid, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if card == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(card))
	}
}

func UpdateCard(c *gin.Context) {

	uri, card, err := validations.UpdateCard(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.UpdateCard(c, uri.ProfileId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	modified, matched, err := card.Update(&uri.CardOid, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if modified == 0 && matched == 0 {
		c.Status(http.StatusNotFound)
	} else if modified == 0 && matched == 1 {
		c.Status(http.StatusNotModified)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(card))
	}
}

func DeleteCards(c *gin.Context) {
	uri, params, err := validations.DeleteCards(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.DeleteCards(c, uri.ProfileId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	cards := models.Cards{}
	res, err := cards.Delete(params.CardOids, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else if res < len(params.CardOids) {
		c.Status(http.StatusMultiStatus)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func DeleteCard(c *gin.Context) {

	uri, err := validations.DeleteCard(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.DeleteCard(c, uri.ProfileId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	card := &models.Card{}

	res, err := card.Delete(&uri.CardOid, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusNoContent)
	}
}
