package utils

import (
	"errors"
	"net/http"

	"github.com/aws/smithy-go"
	"github.com/gin-gonic/gin"
)

const (
	RESPONSE_ERROR = "error"
	RESPONSE_DATA  = "data"
)

type Error struct {
	Message         string `json:"message"`
	Code            int    `json:"code"`
	ApplicationCode string `json:"aplicationCode"`
}

func SuccessBody(data interface{}) gin.H {
	response := gin.H{}
	response[RESPONSE_DATA] = data
	return response
}

func ErrorBody(code int, err error) gin.H {
	response := gin.H{}
	var applicationErr string
	var message = http.StatusText(code)
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) {
			applicationErr = ae.ErrorCode()
			message = ae.ErrorMessage()
		} else {
			message = err.Error()
		}
	}
	response[RESPONSE_ERROR] = Error{Message: message, Code: code, ApplicationCode: applicationErr}
	return response
}
