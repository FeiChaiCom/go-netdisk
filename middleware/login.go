package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	cfg "go-netdisk/config"
	"log"
	"net/http"
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
		cURL := strings.Join([]string{"http://", c.Request.Host, cfg.ENV.Login.SubPath, c.Request.RequestURI}, "")
		referer := c.Request.Header["Referer"]
		if len(referer) > 0 {
			cURL = referer[0]
		}

		redirectURL := cfg.ENV.Login.LoginURL + "?c_url=" + cURL

		// Redirect to pop up window
		if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":    http.StatusUnauthorized,
				"login_url": redirectURL,
			})
			return
		}

		// Redirect to login page
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
	log.Printf("url: %s?ticket=%s, getInfo: %#v\n", cfg.ENV.Login.UserInfoURL, ticket, getInfo)

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
