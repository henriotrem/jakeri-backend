package validations

import (
	"github.com/gin-gonic/gin"
)

type AddSessionBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AddSession(c *gin.Context) (body AddSessionBody, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}

type UpdateSessionBody struct {
	RefreshToken string `json:"refreshtoken" binding:"required"`
}

func UpdateSession(c *gin.Context) (body UpdateSessionBody, err error) {
	if err == nil {
		err = c.ShouldBindJSON(&body)
	}
	return
}
