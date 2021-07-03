package middleware

import (
	"errors"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/models/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	ErrInvalidToken  = errors.New("invalid token")
	ErrLoginRequired = errors.New("login failed")
)

func GetTokenUser(c *gin.Context) (*TokenUser, error) {
	claims, exist := c.Get("claims")
	if !exist {
		return nil, errors.New("claims not exist")
	}

	myClaim, ok := claims.(*MyClaims)
	if !ok {
		return nil, errors.New("parse myclaim error")
	}

	return &myClaim.TokenUser, nil
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

	// log.Printf("%#v, %s", token, err)
	if err != nil {
		return nil, err
	}

	if v, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return v, nil
	}

	return nil, ErrInvalidToken
}

// Login middleware for user auth required apis
func JWTLoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from header or cookie
		// token := c.GetHeader("X-TOKEN")
		token, err := c.Cookie(cfg.AuthCookieName)
		if token == "" || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "login required, token not exist",
			})
			return
		}

		j := JWT{SecretKey: []byte(cfg.JwtSecretKey)}
		claims, err := j.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "login required, token not valid",
			})
			return
		}

		if _, err := db.FindUserByName(claims.Username); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "login required, user not found",
			})
			return
		}

		// Add claims of user info to context
		c.Set("claims", claims)

		// Add for shortcut
		c.Set("UUID", claims.TokenUser.UUID)
		c.Set("username", claims.TokenUser.Username)
		c.Next()
	}
}
