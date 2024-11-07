package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ntquang/ecommerce/global"
)

type PayLoadClams struct {
	jwt.MapClaims
}

func generateTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.API_SERCERT_KEY))
}

func CreateToken(uuidToken string) (string, error) {
	//1. set time expiration
	timeEx := global.Config.JWT.JWT_EXPIRATION
	if timeEx == "" {
		timeEx = "1h"
	}

	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}
	now := time.Now()
	expiresAt := now.Add(expiration)
	return generateTokenJWT(&PayLoadClams{
		MapClaims: jwt.MapClaims{
			"sub": uuidToken,
			"iss": "shopdevgo",
			"jti": uuid.New().String(),
			"exp": expiresAt,
			"iat": time.Now().Unix(),
		},
	})
}
