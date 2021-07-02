package middleware

import (
	"errors"
	"github.com/gaomugong/go-netdisk/models/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

type JWT struct {
	SecretKey []byte
}

type MyClaims struct {
	jwt.StandardClaims
	Username string
	Password string
}

func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(j.SecretKey)
}

func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("sign method not support")
		}
		return j.SecretKey, nil
	})

	log.Printf("%#v, %s", token, err)
	if err != nil {
		return nil, err
	}

	if v, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return v, nil
	}

	return nil, errors.New("invalid token")
}

func JWTLoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-TOKEN")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token not exist",
			})
			return
		}

		j := JWT{SecretKey: []byte("feichaicom")}
		claims, err := j.ParseToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token not valid",
			})
			return
		}

		if _, err := db.FindUserByName(claims.Username); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "username not valid",
			})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
