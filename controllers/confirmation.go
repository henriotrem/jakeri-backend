package controllers

import (
	"jakeri-backend/middleware"
	"jakeri-backend/utils"
	"jakeri-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddConfirmation(c *gin.Context) {

	confirmation, err := validations.AddConfirmation(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	if data, err := middleware.CognitoConfirmRegistration(confirmation); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, err))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(data))
	}
}

func NewConfirmation(c *gin.Context) {

	confirmation, err := validations.NewConfirmation(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	if data, err := middleware.CognitoResendConfirmationCode(confirmation); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, err))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(data))
	}
}
