package validations

import (
	"jakeri-backend/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddUsers(c *gin.Context) (body models.Users, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type GetUsersHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetUsersParams struct {
	IDs     string `form:"ids"`
	UserIDs []string
}

func GetUsers(c *gin.Context) (header GetUsersHeader, params GetUsersParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil && len(params.IDs) > 0 {
		params.UserIDs = strings.Split(params.IDs, ",")
	}
	return
}

type GetUserHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetUserUri struct {
	UserId string `uri:"userId" binding:"required"`
}

func GetUser(c *gin.Context) (header GetUserHeader, uri GetUserUri, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	return
}

type UpdateUserUri struct {
	UserId string `uri:"userId" binding:"required"`
}

func UpdateUser(c *gin.Context) (uri UpdateUserUri, body models.User, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type DeleteUsersParams struct {
	IDs     string `form:"ids" binding:"required"`
	UserIDs []string
}

func DeleteUsers(c *gin.Context) (params DeleteUsersParams, err error) {
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil && len(params.IDs) > 0 {
		params.UserIDs = strings.Split(params.IDs, ",")
	}
	return
}

type DeleteUserUri struct {
	UserId string `uri:"userId" binding:"required"`
}

func DeleteUser(c *gin.Context) (uri DeleteUserUri, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	return
}

type UpdatePasswordUri struct {
	UserId string `uri:"userId" binding:"required"`
}
type UpdatePasswordBody struct {
	New      string `json:"new" binding:"required"`
	Previous string `json:"previous" binding:"required"`
}

func UpdatePassword(c *gin.Context) (uri UpdatePasswordUri, body UpdatePasswordBody, err error) {
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}
