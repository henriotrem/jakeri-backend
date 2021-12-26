package controllers

import (
	"jakeri-backend/authorizations"
	"jakeri-backend/middleware"
	"jakeri-backend/models"
	"jakeri-backend/utils"
	"jakeri-backend/validations"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProfiles(c *gin.Context) {

	uri, profiles, err := validations.AddProfiles(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.AddProfiles(c, &uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	cognitoProfiles := copyProfilesAndLoadUsers(profiles)

	if err = middleware.CognitoAddProfiles(uri.GroupId, cognitoProfiles); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	ids, err := profiles.Add(&uri.GroupOid, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(ids) == 0 {
		c.Status(http.StatusConflict)
	} else if len(ids) < len(profiles) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(ids))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(ids))
	}
}

func GetProfiles(c *gin.Context) {

	header, uri, params, err := validations.GetProfiles(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.GetProfiles(c, &uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	profiles := models.Profiles{}
	err = profiles.Get(&uri.GroupOid, params.ProfileOids, nil, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(profiles) == 0 && len(params.ProfileOids) > 0 {
		c.Status(http.StatusNotFound)
	} else if len(profiles) < len(params.ProfileOids) {
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

	err = authorizations.GetProfile(c, &uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	profile := &models.Profile{}
	err = profile.Get(&uri.GroupOid, &uri.ProfileOid, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if profile == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(profile))
	}
}

func UpdateProfile(c *gin.Context) {

	uri, profile, err := validations.UpdateProfile(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.UpdateProfile(c, &uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	modified, matched, err := profile.Update(&uri.GroupOid, &uri.ProfileOid, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if modified == 0 && matched == 0 {
		c.Status(http.StatusNotFound)
	} else if modified == 0 && matched == 1 {
		c.Status(http.StatusNotModified)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(profile))
	}
}

func DeleteProfiles(c *gin.Context) {

	uri, params, err := validations.DeleteProfiles(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.DeleteProfiles(c, &uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	profiles := models.Profiles{}
	embed := map[string]interface{}{
		"user": map[string]interface{}{},
	}
	profiles.Get(&uri.GroupOid, params.ProfileOids, nil, embed)

	if err = middleware.CognitoDeleteProfiles(uri.GroupId, profiles); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	res, err := profiles.Delete(&uri.GroupOid, params.ProfileOids)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else if res < len(params.ProfileOids) {
		c.Status(http.StatusMultiStatus)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func DeleteProfile(c *gin.Context) {

	uri, err := validations.DeleteProfile(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.DeleteProfile(c, &uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	profile := &models.Profile{}
	embed := map[string]interface{}{
		"user": map[string]interface{}{},
	}
	profile.Get(&uri.GroupOid, &uri.ProfileOid, embed)

	if _, err = middleware.CognitoDeleteProfile(uri.GroupId, profile); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	res, err := profile.Delete(&uri.GroupOid, &uri.ProfileOid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func copyProfilesAndLoadUsers(profiles models.Profiles) models.Profiles {

	newProfiles := models.Profiles{}

	for _, profile := range profiles {
		id := *profile.User.ID
		newProfile := &models.Profile{User: &models.User{ID: &id}}
		newProfile.User.Get(*newProfile.User.ID, nil)
		newProfiles = append(newProfiles, *newProfile)
	}

	return newProfiles
}
