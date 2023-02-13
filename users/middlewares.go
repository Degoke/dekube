package users

import (
	"net/http"
	"strings"

	"github.com/Degoke/dekube/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
)

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}

	return tok, nil
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorizartion"},
	Filter: stripBearerPrefixFromTokenString,
}

func UpdateContextUserModel(c *gin.Context, user_id uint) {
	var userModel UserModel
	if user_id != 0 {
		db := common.GetDB()
		db.First(&userModel, user_id)
	}
	c.Set("user_id", user_id)
	c.Set("user_model", userModel)
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		token, err := request.ParseFromRequest(c.Request, AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			secret := common.GetENV("NB_SECRET")
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user_id := uint(claims["id"].(float64))
			UpdateContextUserModel(c, user_id)
		}
	}
}