package login

import (
	"fmt"
	cfg "github.com/gaomugong/go-netdisk/config"
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
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err.Error(),
		})
		return
	}

	// Verify username & password and login
	u := &db.User{Username: p.Username, Password: p.Password}
	validUser, err := db.Login(u)
	fmt.Printf("%#v, %s", validUser, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"error":   err.Error(),
			"message": "invalid user or password",
		})
		return
	}

	// Make token response with user claim
	nowTime := time.Now()
	expiredTime := nowTime.Add(time.Hour * 24).Unix()
	j := &middleware.JWT{SecretKey: []byte(cfg.JwtSecretKey)}
	claims := middleware.MyClaims{
		TokenUser: middleware.TokenUser{
			UUID:     validUser.UUID.String(),
			Username: validUser.Username,
			Password: validUser.Password,
		},
		StandardClaims: jwt.StandardClaims{
			Issuer:    cfg.JwtIssuer,
			ExpiresAt: expiredTime,
		},
	}

	token, _ := j.CreateToken(claims)
	// c.Header("X-TOKEN", token)
	c.SetCookie(cfg.AuthCookieName, token, 60*60*24, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"user":        validUser,
		"accessToken": token,
		"ExpiresAt":   claims.StandardClaims.ExpiresAt * 1000,
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
