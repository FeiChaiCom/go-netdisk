package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	cfg "go-netdisk/config"
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
		// if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		// 	cURL := strings.Join([]string{"http://", c.Request.Host, os.Getenv(cfg.ENV.Login.SubPath), "account/login_success"}, "")
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 		"status":    http.StatusUnauthorized,
		// 		"login_url": cfg.ENV.Login.LoginURL + "?c_url=" + cURL,
		// 		"width":     460,
		// 		"height":    490,
		// 	})
		// 	return
		// }

		// Redirect to login page
		appSubPath := os.Getenv(cfg.ENV.Login.SubPath)
		subPath := appSubPath[:len(appSubPath)-1]
		referURL := strings.Join([]string{"http://", c.Request.Host, subPath, c.Request.RequestURI}, "")
		redirectUrl := strings.Join([]string{subPath, "account/login_page?refer_url=", referURL}, "")

		c.Redirect(http.StatusFound, redirectUrl)
	}

	// verify session and cookie success
	if ticket == session.Get(cfg.ENV.Login.Ticket) {
		c.Next()
		return
	}

	// Verify cookie from remote login server
	getInfo := GetInfo{}
	request := gorequest.New()
	_, _, errs := request.Get(cfg.ENV.Login.UserInfoURL).Query(map[string]string{cfg.ENV.Login.Ticket: ticket}).EndStruct(&getInfo)
	if errs != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"状态返回": "内部错误"})
		return
	}

	// Login success from remote login server
	if getInfo.Ret == 0 {
		uid, err := c.Cookie(cfg.ENV.Login.UID)
		if err == nil {
			session.Set(cfg.ENV.Login.Ticket, ticket)
			session.Set(cfg.ENV.Login.UID, uid)
			session.Save()
			c.Next()
			return
		}
	}

}
