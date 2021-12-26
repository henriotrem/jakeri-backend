package controllers

import (
	"jakeri-backend/authorizations"
	"jakeri-backend/models"
	"jakeri-backend/utils"
	"jakeri-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddReviews(c *gin.Context) {

	reviews, err := validations.AddReviews(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.AddReviews(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	ids, err := reviews.Add(&tokenData.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(ids) == 0 {
		c.Status(http.StatusConflict)
	} else if len(ids) < len(reviews) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(ids))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(ids))
	}
}

func GetReviews(c *gin.Context) {

	header, params, err := validations.GetReviews(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.GetReviews(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	reviews := models.Reviews{}
	err = reviews.Get(params.ReviewsOids, header.Embed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(reviews) == 0 && len(params.ReviewsOids) > 0 {
		c.Status(http.StatusNotFound)
	} else if len(reviews) < len(params.ReviewsOids) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(reviews))
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(reviews))
	}
}

func GetReview(c *gin.Context) {

	header, uri, err := validations.GetReview(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.GetReview(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	news := &models.Review{}
	err = news.Get(&uri.ReviewOid, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if news == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(news))
	}
}
