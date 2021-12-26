package authorizations

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenData struct {
	UserID      string
	Username    string
	GroupIds    map[primitive.ObjectID]bool
	Roles       map[string]bool
	AccessToken string
}

func GetAccessTokenData(c *gin.Context) TokenData {

	tokenMap := GetTokenMap(c)

	tokenData := TokenData{
		UserID:      GetUserIDFromTokenMap(tokenMap),
		Username:    GetUsernameFromTokenMap(tokenMap),
		GroupIds:    GetGroupIdFromTokenMap(tokenMap),
		Roles:       GetRolesFromTokenMap(tokenMap),
		AccessToken: GetAccessToken(c),
	}
	return tokenData
}
func GetAccessToken(c *gin.Context) string {
	var accessToken string
	tmp, ok := c.Get("accessToken")
	if !ok {
		return accessToken
	}
	accessToken, ok = tmp.(string)
	if !ok {
		return accessToken
	}
	return accessToken
}

func GetTokenMap(c *gin.Context) map[string]interface{} {
	var tokenMap map[string]interface{}
	tmp, ok := c.Get("tokenMap")
	if !ok {
		return tokenMap
	}
	tokenMap, ok = tmp.(map[string]interface{})
	if !ok {
		return tokenMap
	}
	return tokenMap
}

func GetUserIDFromTokenMap(tokenMap map[string]interface{}) string {
	var userId string
	tmp, ok := tokenMap["sub"]
	if !ok {
		return userId
	}
	str, ok := tmp.(string)
	if !ok {
		return userId
	}
	userId = str
	return userId
}

func GetUsernameFromTokenMap(tokenMap map[string]interface{}) string {
	var username string
	tmp, ok := tokenMap["username"]
	if !ok {
		return username
	}
	str, ok := tmp.(string)
	if !ok {
		return username
	}
	username = str
	return username
}

func GetGroupIdFromTokenMap(tokenMap map[string]interface{}) map[primitive.ObjectID]bool {
	var groupIDs = make(map[primitive.ObjectID]bool)
	tmp, ok := tokenMap["cognito:groups"]
	if !ok {
		return groupIDs
	}
	groups, ok := tmp.([]interface{})
	if !ok {
		return groupIDs
	}
	for _, group := range groups {
		str, ok := group.(string)
		if !ok {
			continue
		}
		tmp := strings.Split(str, ":")
		if tmp[0] == "group" {
			id, _ := primitive.ObjectIDFromHex(tmp[1])
			groupIDs[id] = true
		}
	}
	return groupIDs
}

func GetRolesFromTokenMap(tokenMap map[string]interface{}) map[string]bool {
	var roles = make(map[string]bool)
	tmp, ok := tokenMap["cognito:groups"]
	if !ok {
		return roles
	}
	groups, ok := tmp.([]interface{})
	if !ok {
		return roles
	}
	for _, group := range groups {
		str, ok := group.(string)
		if !ok {
			continue
		}
		tmp := strings.Split(str, ":")
		if tmp[0] == "role" {
			roles[tmp[1]] = true
		}
	}
	return roles
}
