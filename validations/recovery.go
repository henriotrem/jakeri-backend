package validations

import (
	"github.com/gin-gonic/gin"
)

type AddRecoveryBody struct {
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ConfirmationCode string `json:"confirmationCode" binding:"required"`
}

func AddRecovery(c *gin.Context) (body AddRecoveryBody, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type NewRecoveryBody struct {
	Username string `json:"username" binding:"required"`
}

func NewRecovery(c *gin.Context) (body NewRecoveryBody, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}
