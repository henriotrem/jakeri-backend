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

func AddUsers(c *gin.Context) {

	users, err := validations.AddUsers(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	if err := middleware.CognitoRegister(users); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
		return
	}

	ids, err := users.Add()

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(ids) == 0 {
		c.Status(http.StatusConflict)
	} else if len(ids) < len(users) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(ids))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(ids))
	}
}

func GetUsers(c *gin.Context) {

	header, params, err := validations.GetUsers(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.GetUsers(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	users := models.Users{}
	err = users.Get(params.UserIDs, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(users) == 0 && len(params.UserIDs) > 0 {
		c.Status(http.StatusNotFound)
	} else if len(users) < len(params.UserIDs) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(users))
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(users))
	}
}

func GetUser(c *gin.Context) {

	header, uri, err := validations.GetUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.GetUser(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	user := &models.User{}
	err = user.Get(uri.UserId, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if user == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(user))
	}
}

func UpdateUser(c *gin.Context) {

	uri, user, err := validations.UpdateUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.UpdateUser(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	modified, matched, err := user.Update(&uri.UserId, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if modified == 0 && matched == 0 {
		c.Status(http.StatusNotFound)
	} else if modified == 0 && matched == 1 {
		c.Status(http.StatusNotModified)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(user))
	}
}

func DeleteUsers(c *gin.Context) {

	params, err := validations.DeleteUsers(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.DeleteUsers(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	users := models.Users{}

	if err := users.Get(params.UserIDs, nil); err != nil {
		c.JSON(http.StatusNotFound, utils.SuccessBody(err))
		return
	}
	if err := middleware.CognitoDeleteUsers(users); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
		return
	}

	res, err := users.Delete(params.UserIDs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else if res < len(params.UserIDs) {
		c.Status(http.StatusMultiStatus)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func DeleteUser(c *gin.Context) {

	uri, err := validations.DeleteUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.DeleteUser(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	user := &models.User{}

	if err := user.Get(uri.UserId, nil); err != nil {
		c.JSON(http.StatusNotFound, utils.SuccessBody(err))
		return
	}
	if err := middleware.CognitoDeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
		return
	}

	res, err := user.Delete(uri.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func UpdatePassword(c *gin.Context) {

	uri, password, err := validations.UpdatePassword(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.UpdatePassword(c, uri.UserId)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	if data, err := middleware.CognitoChangePassword(&tokenData.AccessToken, password); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(data))
	}
}
