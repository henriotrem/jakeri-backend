package validations

import (
	"github.com/gin-gonic/gin"
)

type AddConfirmationBody struct {
	ConfirmationCode string `json:"confirmationCode" binding:"required"`
	Username         string `json:"username" binding:"required"`
}

func AddConfirmation(c *gin.Context) (body AddConfirmationBody, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type NewConfirmationBody struct {
	Username string `json:"username" binding:"required"`
}

func NewConfirmation(c *gin.Context) (body NewConfirmationBody, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}
