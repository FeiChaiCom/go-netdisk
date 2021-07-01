package middleware

import (
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SecretKey []byte
}

type AuthClaims struct {
	jwt.StandardClaims
	Username string
	Password string
}

func (j *JWT) CreateToken(claims AuthClaims) (string, error) {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(j.SecretKey)
}
