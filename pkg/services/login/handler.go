package login

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-netdisk/pkg/db/models"
	"go-netdisk/pkg/settings"
	"go-netdisk/pkg/utils"

	"go-netdisk/pkg/db"
	"go-netdisk/pkg/middlewares"
	"net/http"
	"time"
)

// After login, update login ip and time
func AfterLoginSuccess(c *gin.Context, u *models.User) error {
	// u.UpdateTime = time.Now()
	// u.LastTime = time.Now()
	// u.LastIP = c.GetHeader("X-Forwarded-For")
	// return db.DB.Save(u).Error
	return db.DB.Model(&models.User{}).Where("username = ?", u.Username).Updates(models.User{
		UpdateTime: time.Now(),
		LastTime:   time.Now(),
		LastIP:     c.GetHeader("X-Forwarded-For"),
	}).Error
}

// Fake logout to clean cookie
func JwtLogoutHandler(c *gin.Context) {
	// Delete cookie
	c.SetCookie(settings.ENV.JWT.AuthCookieName, "", -1, "/", "", false, true)
	utils.OkOnly(c)
}

// curl http://localhost:5000/api/account/login/ -X POST -d '{"username": "miya", "password": "miya.12345"}'
func JwtLoginHandler(c *gin.Context) {
	var p *models.LoginParam
	if err := c.ShouldBind(&p); err != nil {
		utils.Error(c, err)
		return
	}

	// Verify username & password and login
	u, err := models.Login(p)
	if err != nil {
		utils.Error(c, err)
		return
	}

	// Make token response with user claim
	j := &middlewares.JWT{SecretKey: []byte(settings.ENV.JWT.SecretKey)}
	claims := middlewares.MyClaims{
		TokenUser: middlewares.TokenUser{
			UUID:     u.UUID.String(),
			Username: u.Username,
			Password: u.Password,
		},
		StandardClaims: jwt.StandardClaims{
			Issuer:    settings.ENV.JWT.Issuer,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token, _ := j.CreateToken(claims)

	// c.Header("X-TOKEN", token)
	c.SetCookie(settings.ENV.JWT.AuthCookieName, token, 60*60*24, "/", "", false, true)

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
	var r *models.RegisterParam

	if err := c.ShouldBindJSON(&r); err != nil {
		utils.Error(c, err)
	}

	user, err := models.Register(r)
	if err != nil {
		utils.Error(c, err)
		return
	}

	utils.Ok(c, user)
}

func LoginSuccessHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login_success.html", gin.H{})
}

func LoginPageHandler(c *gin.Context) {
	referURL := c.Query("refer_url")
	c.HTML(http.StatusUnauthorized, "login_page.html", gin.H{
		"refer_url":         referURL,
		"static_url":        "/static/",
		"remote_static_url": "/static/",
	})
}
