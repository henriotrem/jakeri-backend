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

func AddGroups(c *gin.Context) {

	groups, err := validations.AddGroups(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.AddGroups(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	ids, err := groups.Add(&tokenData.UserID)

	if err := middleware.CognitoCreateGroups(groups, ids); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(ids) == 0 {
		c.Status(http.StatusConflict)
	} else if len(ids) < len(groups) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(ids))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(ids))
	}
}

func GetGroups(c *gin.Context) {

	header, params, err := validations.GetGroups(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.GetGroups(c, &params.GroupOids)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	groups := models.Groups{}
	err = groups.Get(params.GroupOids, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(groups) == 0 && len(params.GroupOids) > 0 {
		c.Status(http.StatusNotFound)
	} else if len(groups) < len(params.GroupOids) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(groups))
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(groups))
	}
}

func GetGroup(c *gin.Context) {

	header, uri, err := validations.GetGroup(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.GetGroup(c, &uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	group := &models.Group{}
	err = group.Get(&uri.GroupOid, header.Embed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if group == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(group))
	}
}

func UpdateGroup(c *gin.Context) {

	uri, group, err := validations.UpdateGroup(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.UpdateGroup(c, &uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	modified, matched, err := group.Update(&uri.GroupOid, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if modified == 0 && matched == 0 {
		c.Status(http.StatusNotFound)
	} else if modified == 0 && matched == 1 {
		c.Status(http.StatusNotModified)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(group))
	}
}

func DeleteGroups(c *gin.Context) {

	params, err := validations.DeleteGroups(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.DeleteGroups(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	groups := models.Groups{}

	if err := groups.Get(params.GroupOids, nil); err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorBody(http.StatusNotFound, err))
		return
	}
	if err := middleware.CognitoDeleteGroups(groups); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
		return
	}

	res, err := groups.Delete(params.GroupOids)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else if res < len(params.GroupOids) {
		c.Status(http.StatusMultiStatus)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func DeleteGroup(c *gin.Context) {

	uri, err := validations.DeleteGroup(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.DeleteGroup(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	group := &models.Group{}

	if err := group.Get(&uri.GroupOid, nil); err != nil {
		c.JSON(http.StatusNotFound, utils.SuccessBody(err))
		return
	}
	if _, err := middleware.CognitoDeleteGroup(group); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
		return
	}

	res, err := group.Delete(&uri.GroupOid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusNoContent)
	}
}
