package controllers

import (
	"jakeri-backend/middleware"
	"jakeri-backend/utils"
	"jakeri-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRecovery(c *gin.Context) {

	recovery, err := validations.AddRecovery(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	if data, err := middleware.CognitoConfirmForgotPassword(recovery); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, err))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(data))
	}
}

func NewRecovery(c *gin.Context) {

	recovery, err := validations.NewRecovery(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	if data, err := middleware.CognitoForgotPassword(recovery); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, err))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(data))
	}
}
