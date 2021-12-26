package middleware

import (
	"context"
	"errors"
	"jakeri-backend/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

var publicKeySet jwk.Set

func init() {
	publicKeySet, _ = jwk.ParseString(os.Getenv("COGNITO_PUBLIC_KEY"))
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
	}
}

func validateToken(c *gin.Context) {
	var err error
	tokenString := c.GetHeader("Authorization")[len("Bearer")+1:]

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, nil))
		return
	}

	var tokenMap map[string]interface{}
	if tokenMap, err = getTokenMap(tokenString); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, nil))
		return
	}

	c.Set("accessToken", tokenString)
	c.Set("tokenMap", tokenMap)

	if err = checkExpiry(tokenMap); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, nil))
		return
	}

	if err = validateIssuer(tokenMap); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, nil))
		return
	}

	if err = validateClient(tokenMap); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, nil))
		return
	}

	if err = validateUse(tokenMap); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorBody(http.StatusUnauthorized, nil))
		return
	}
}

func getTokenMap(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseString(tokenString, jwt.WithKeySet(publicKeySet))
	if err != nil {
		return nil, err
	}

	// convert the JWT to map to access all required fields
	tokenMap, err := token.AsMap(context.Background())
	if err != nil {
		return nil, err
	}
	return tokenMap, nil
}

//check expiry
func checkExpiry(tokenMap map[string]interface{}) error {
	expiration := tokenMap["exp"].(time.Time)
	// expiration := tokenMap["exp"].(time.Time).Add(-time.Minute * 55) // make it expire after 5 minutes for test purposes
	if time.Now().After(expiration) {
		return errors.New("error: token has expired")
	}
	return nil
}

// validate issuer
func validateIssuer(tokenMap map[string]interface{}) error {
	issuer := tokenMap["iss"].(string)
	if issuer != "https://cognito-idp."+Region+".amazonaws.com/"+UserPoolID {
		return errors.New("token issuer is invalid")
	}
	return nil
}

// validate audience
func validateClient(tokenMap map[string]interface{}) error {
	clientID := tokenMap["client_id"].(string)
	if clientID != ClientID {
		return errors.New("token audience is invalid")
	}
	return nil
}

// validate use
func validateUse(tokenMap map[string]interface{}) error {
	use := tokenMap["token_use"].(string)
	if use != "access" {
		return errors.New("token use is invalid")
	}
	return nil
}
