package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/services/demo"
	"go-netdisk/services/login"
	"go-netdisk/services/matter"
	"go-netdisk/services/monitor"
	"go-netdisk/services/permission"
	"go-netdisk/services/preference"
	"go-netdisk/services/user"
	"go-netdisk/settings"
	"net/http"
)

type Register func(rg *gin.RouterGroup)

var registers = []Register{
	login.RegisterLoginGroup,
	user.RegisterUserGroup,
	matter.RegisterMatterGroup,
	preference.RegisterPreferenceGroup,
	permission.RegisterPermissionGroup,
	monitor.RegisterMonitorGroup,
	demo.RegisterTestGroup,
}

func InitRouter(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	for _, register := range registers {
		register(apiGroup)
	}
	// Init template and static files serve router
	initTemplateRouter(engine)
}

func initTemplateRouter(engine *gin.Engine) {
	// Load index html
	engine.LoadHTMLGlob(settings.TemplateDirPattern)
	engine.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "feichai",
		})
	})

	// Serve static files
	engine.Static(settings.StaticURL, settings.StaticDir)
	engine.StaticFile("/favicon.ico", fmt.Sprintf("%s/favicon.ico", settings.StaticDir))

	// Serve media files
	engine.StaticFS(settings.MediaURL, http.Dir(settings.ENV.MediaDir))
}
