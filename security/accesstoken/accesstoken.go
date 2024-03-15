package accesstoken

import (
	"github.com/golang-jwt/jwt"
	"shop-backend/config"
)

var tokenSecret = []byte(config.Cfg.AccessToken)

func GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	return token.SignedString(tokenSecret)
}
