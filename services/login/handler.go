package login

import (
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/middleware"
	"github.com/gaomugong/go-netdisk/models/db"
	R "github.com/gaomugong/go-netdisk/render"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

// After login, update login ip and time
func AfterLoginSuccess(c *gin.Context, u *db.User) error {
	// u.UpdateTime = time.Now()
	// u.LastTime = time.Now()
	// u.LastIP = c.GetHeader("X-Forwarded-For")
	// return cfg.DB.Save(u).Error
	return cfg.DB.Model(&db.User{}).Where("username = ?", u.Username).Updates(db.User{
		UpdateTime: time.Now(),
		LastTime:   time.Now(),
		LastIP:     c.GetHeader("X-Forwarded-For"),
	}).Error
}

// curl http://localhost:5000/api/account/login/ -X POST -d '{"username": "miya", "password": "miya.12345"}'
func JwtLoginHandler(c *gin.Context) {
	var p *db.LoginParam
	if err := c.ShouldBind(&p); err != nil {
		R.Error(c, err)
		return
	}

	// Verify username & password and login
	u, err := db.Login(p)
	if err != nil {
		R.Error(c, err)
		return
	}

	// Make token response with user claim
	j := &middleware.JWT{SecretKey: []byte(cfg.JwtSecretKey)}
	claims := middleware.MyClaims{
		TokenUser: middleware.TokenUser{
			UUID:     u.UUID.String(),
			Username: u.Username,
			Password: u.Password,
		},
		StandardClaims: jwt.StandardClaims{
			Issuer:    cfg.JwtIssuer,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token, _ := j.CreateToken(claims)

	// c.Header("X-TOKEN", token)
	c.SetCookie(cfg.AuthCookieName, token, 60*60*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user":        u,
		"accessToken": token,
		"ExpiresAt":   claims.StandardClaims.ExpiresAt * 1000,
	})

	// TODO: Login hook
	_ = AfterLoginSuccess(c, u)
}

// curl http://localhost:5000/api/account/register/ -X POST -d '{"username": "miya", "password": "miya.12345"}'
func RegisterHandler(c *gin.Context) {
	var r *db.RegisterParam

	if err := c.ShouldBindJSON(&r); err != nil {
		R.Error(c, err)
	}

	user, err := db.Register(r)
	if err != nil {
		R.Error(c, err)
		return
	}

	R.Ok(c, user)
}
