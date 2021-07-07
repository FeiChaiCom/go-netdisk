package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	cfg "go-netdisk/config"
	"go-netdisk/models/db"
	"net/http"
	"strings"
)

type UserInfo struct {
	Message string           `json:"msg"`
	Data    UserFullInfoData `json:"data"`
	Ret     int              `json:"ret"`
}
type UserInfoData struct {
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

type UserFullInfoData struct {
	UserInfoData
	DeptName  string `json:"dept_name"`
	GroupName string `json:"group_name"`
	PostName  string `json:"post_name"`
}

func LoginRequired(c *gin.Context) {
	session := sessions.Default(c)
	ticket, err := c.Cookie(cfg.ENV.Login.Ticket)

	// Redirect user to login first
	if err != nil {
		cURL := strings.Join([]string{"http://", c.Request.Host, cfg.ENV.Login.SubPath, c.Request.RequestURI}, "")
		if referer := c.Request.Header["Referer"]; len(referer) > 0 {
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
		c.Redirect(http.StatusFound, redirectURL)
		return
	}

	// verify session and cookie success
	if ticket == session.Get(cfg.ENV.Login.Ticket) {
		// Inject username to context
		uid, _ := c.Cookie(cfg.ENV.Login.UID)
		c.Set("username", uid)
		c.Next()
		return
	}

	// Verify cookie from remote login server
	var userInfo UserInfo
	_, _, errs := gorequest.New().Get(cfg.ENV.Login.UserInfoURL).Query(map[string]string{
		cfg.ENV.Login.Ticket: ticket,
	}).EndStruct(&userInfo)
	if errs != nil || userInfo.Ret != 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "server error, query user info failed"})
		return
	}

	if _, err := db.GetOrCreateUser(userInfo.Data.Username); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "server error, create user info failed"})
		return
	}
	// log.Printf("url: %s?ticket=%s, userInfo: %#v\n", cfg.ENV.Login.UserInfoURL, ticket, userInfo)

	// Login success from remote login server
	uid, err := c.Cookie(cfg.ENV.Login.UID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "login required, user not found"})
		return
	}
	session.Set(cfg.ENV.Login.Ticket, ticket)
	session.Set(cfg.ENV.Login.UID, uid)
	session.Save()

	// Inject username to context
	c.Set("username", uid)
	c.Next()
}
