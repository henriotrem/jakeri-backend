package controllers

import (
	"jakeri-backend/authorizations"
	"jakeri-backend/models"
	"jakeri-backend/utils"
	"jakeri-backend/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPosts(c *gin.Context) {

	posts, err := validations.AddPosts(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.AddPosts(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	ids, err := posts.Add(&tokenData.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(ids) == 0 {
		c.Status(http.StatusConflict)
	} else if len(ids) < len(posts) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(ids))
	} else {
		c.JSON(http.StatusCreated, utils.SuccessBody(ids))
	}
}

func GetPosts(c *gin.Context) {

	header, params, err := validations.GetPosts(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	posts := models.Posts{}
	err = posts.Get(params.PostOids, header.Embed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if len(posts) == 0 && len(params.PostOids) > 0 {
		c.Status(http.StatusNotFound)
	} else if len(posts) < len(params.PostOids) {
		c.JSON(http.StatusMultiStatus, utils.SuccessBody(posts))
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(posts))
	}
}

func GetPost(c *gin.Context) {

	uri, err := validations.GetPost(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	post := &models.Post{}
	err = post.Get(&uri.PostOid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if post == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(post))
	}
}

func UpdatePost(c *gin.Context) {

	uri, post, err := validations.UpdatePost(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	tokenData, err := authorizations.UpdatePost(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	modified, matched, err := post.Update(&uri.PostOid, &tokenData.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if modified == 0 && matched == 0 {
		c.Status(http.StatusNotFound)
	} else if modified == 0 && matched == 1 {
		c.Status(http.StatusNotModified)
	} else {
		c.JSON(http.StatusOK, utils.SuccessBody(post))
	}
}

func DeletePosts(c *gin.Context) {
	params, err := validations.DeletePosts(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.DeletePosts(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	posts := models.Posts{}
	res, err := posts.Delete(params.PostOids)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else if res < len(params.PostOids) {
		c.Status(http.StatusMultiStatus)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func DeletePost(c *gin.Context) {

	uri, err := validations.DeletePost(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorBody(http.StatusBadRequest, err))
		return
	}

	err = authorizations.DeletePost(c)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorBody(http.StatusForbidden, err))
		return
	}

	post := &models.Post{}

	res, err := post.Delete(&uri.PostOid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorBody(http.StatusInternalServerError, err))
	} else if res == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.Status(http.StatusNoContent)
	}
}
