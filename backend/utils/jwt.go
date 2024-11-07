package utils

import (
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(claims jwt.Claims, method jwt.SigningMethod, jwtSecret string) (string, error) {

	token := jwt.NewWithClaims(method, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
