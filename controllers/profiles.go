package controllers

import (
	"jakeri-backend/models"
	"jakeri-backend/utils"
	"jakeri-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfiles(c *gin.Context) {

	header, params, err := validations.GetProfiles(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	profiles := models.Profiles{}
	err = profiles.Get(params.ProfileIDs, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(profiles) == 0 && len(params.ProfileIDs) > 0 {
		c.Status(http.StatusNotFound)
	} else if len(profiles) < len(params.ProfileIDs) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(profiles))
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(profiles))
	}
}

func GetProfile(c *gin.Context) {

	header, uri, err := validations.GetProfile(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	profile := &models.Profile{}
	err = profile.Get(uri.ProfileId, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if profile == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(profile))
	}
}
