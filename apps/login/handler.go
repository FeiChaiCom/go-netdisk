package login

import (
	"github.com/gaomugong/go-netdisk/middleware"
	"github.com/gaomugong/go-netdisk/models/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type loginParam struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// curl http://localhost:5000/api/account/login/ -X POST -d '{"username": "miya", "password": "miya.12345"}'
func AuthHandler(c *gin.Context) {
	var p loginParam
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	// Verify username & password and return jwt token as header
	u := &db.User{Username: p.Username, Password: p.Password}

	validUser, err := db.Login(u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	// Fetch user by username
	j := &middleware.JWT{SecretKey: []byte("feichaicom")}
	claims := middleware.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "feichai",
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Add(1).Unix(),
		},
		Username: validUser.Username,
		Password: validUser.Password,
	}

	token, _ := j.CreateToken(claims)
	c.Header("x-token", token)

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data": gin.H{
			"User":      validUser,
			"Token":     token,
			"ExpiresAt": claims.StandardClaims.ExpiresAt * 1000,
		},
		"message": "login success",
	})
}

type registerParam struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// curl http://localhost:5000/api/account/register/ -X POST -d '{"username": "miya", "password": "miya.12345"}'
func RegisterHandler(c *gin.Context) {
	var r registerParam

	_ = c.ShouldBindJSON(&r)
	user := db.User{Username: r.Username, Password: r.Password}
	newUser, err := db.Register(user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"data":   newUser,
	})
}
