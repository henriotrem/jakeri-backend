package controllers

import (
	"jakeri-backend/middleware"
	"jakeri-backend/utils"
	"jakeri-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSession(c *gin.Context) {

	session, err := validations.AddSession(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	if data, err := middleware.CognitoLogin(session); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, err))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(data))
	}
}

func UpdateSession(c *gin.Context) {

	session, err := validations.UpdateSession(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	if data, err := middleware.CognitoRefreshLogin(session); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, err))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(data))
	}
}
