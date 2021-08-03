package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/services/demo"
	"go-netdisk/pkg/services/login"
	"go-netdisk/pkg/services/matter"
	"go-netdisk/pkg/services/monitor"
	"go-netdisk/pkg/services/permission"
	"go-netdisk/pkg/services/preference"
	"go-netdisk/pkg/services/user"
	"go-netdisk/pkg/settings"
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
