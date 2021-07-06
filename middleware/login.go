package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	cfg "go-netdisk/config"
	"log"
	"net/http"
	"os"
	"strings"
)

type GetInfo struct {
	Message string   `json:"msg"`
	Data    UserInfo `json:"data"`
	Ret     int      `json:"ret"`
}
type UserInfo struct {
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

func LoginRequired(c *gin.Context) {
	session := sessions.Default(c)
	ticket, err := c.Cookie(cfg.ENV.Login.Ticket)

	// Redirect user to login first
	if err != nil {
		// Redirect to pop up window
		if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
			cURL := strings.Join([]string{"http://", c.Request.Host, os.Getenv(cfg.ENV.Login.SubPath), "account/login_success"}, "")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":    http.StatusUnauthorized,
				"login_url": cfg.ENV.Login.LoginURL + "?c_url=" + cURL,
				"width":     460,
				"height":    490,
			})
			return
		}

		// Redirect to login page
		//appSubPath := os.Getenv(cfg.ENV.Login.SubPath)
		//subPath := appSubPath[:len(appSubPath)-1]
		//referURL := strings.Join([]string{"http://", c.Request.Host, subPath, c.Request.RequestURI}, "")
		//redirectURL := strings.Join([]string{subPath, "account/login_page?refer_url=", referURL}, "")

		cURL := strings.Join([]string{"http://", c.Request.Host, os.Getenv(cfg.ENV.Login.SubPath), "/"}, "")
		redirectURL := cfg.ENV.Login.LoginURL + "?c_url=" + cURL
		log.Printf("Redirect to login page: %s\n", redirectURL)
		c.Redirect(http.StatusFound, redirectURL)
		return
	}

	// verify session and cookie success
	if ticket == session.Get(cfg.ENV.Login.Ticket) {
		c.Next()
		return
	}

	// Verify cookie from remote login server
	var getInfo GetInfo
	_, _, errs := gorequest.New().Get(cfg.ENV.Login.UserInfoURL).Query(map[string]string{
		cfg.ENV.Login.Ticket: ticket,
	}).EndStruct(&getInfo)
	if errs != nil || getInfo.Ret != 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "server error, query user info failed"})
		return
	}
	log.Printf("getInfo: %#v\n", getInfo)

	// Login success from remote login server
	uid, err := c.Cookie(cfg.ENV.Login.UID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "login required, user not found",
		})
		return
	}

	log.Printf("login sucess: uid=%s, ticket=%s\n", uid, ticket)
	session.Set(cfg.ENV.Login.Ticket, ticket)
	session.Set(cfg.ENV.Login.UID, uid)
	session.Save()
	c.Next()

}
