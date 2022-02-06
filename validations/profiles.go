package validations

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type GetProfilesHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetProfilesParams struct {
	IDs        string `form:"ids"`
	ProfileIDs []string
}

func GetProfiles(c *gin.Context) (header GetProfilesHeader, params GetProfilesParams, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindQuery(&params)
	}
	if err == nil && len(params.IDs) > 0 {
		params.ProfileIDs = strings.Split(params.IDs, ",")
	}
	return
}

type GetProfileHeader struct {
	Embed map[string]interface{} `header:"Embed"`
}
type GetProfileUri struct {
	ProfileId string `uri:"profileId" binding:"required"`
}

func GetProfile(c *gin.Context) (header GetProfileHeader, uri GetProfileUri, err error) {
	if err == nil {
		err = c.ShouldBindHeader(&header)
	}
	if err == nil {
		err = c.ShouldBindUri(&uri)
	}
	return
}
