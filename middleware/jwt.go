package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	cfg "go-netdisk/config"
	"go-netdisk/models/db"
	"net/http"
)

type JWT struct {
	SecretKey []byte
}

type TokenUser struct {
	UUID     string
	Username string
	Password string
}

type MyClaims struct {
	jwt.StandardClaims
	TokenUser
}

var (
	ErrInvalidToken = errors.New("invalid token")
)

// Login middleware for user auth required apis
func JWTLoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from header or cookie
		// token := c.GetHeader("X-TOKEN")
		token, err := c.Cookie(cfg.ENV.JWT.AuthCookieName)
		if token == "" || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "login required, token not exist",
			})
			return
		}

		j := JWT{SecretKey: []byte(cfg.ENV.JWT.SecretKey)}
		claims, err := j.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "login required, token not valid",
			})
			return
		}

		// user, err := db.GetUserByUUID(claims.UUID)
		user, err := db.GetUserByName(claims.Username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "login required, user not found",
			})
			return
		}

		// Add user info to context
		c.Set("user", user)
		c.Set("UUID", claims.TokenUser.UUID)
		c.Set("username", claims.TokenUser.Username)
		c.Next()
	}
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

	if err != nil {
		return nil, err
	}

	if v, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return v, nil
	}

	return nil, ErrInvalidToken
}
