package services

import (
	"fmt"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gaomugong/go-netdisk/middleware"
	"github.com/gaomugong/go-netdisk/services/demo"
	"github.com/gaomugong/go-netdisk/services/login"
	"github.com/gaomugong/go-netdisk/services/matter"
	"github.com/gaomugong/go-netdisk/services/monitor"
	"github.com/gaomugong/go-netdisk/services/preference"
	"github.com/gaomugong/go-netdisk/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Register func(rg *gin.RouterGroup)

var registers = []Register{
	login.RegisterLoginGroup,
	user.RegisterUserGroup,
	matter.RegisterMatterGroup,
	preference.RegisterPreferenceGroup,
	monitor.RegisterMonitorGroup,
	demo.RegisterTestGroup,
}

func InitAPIRouter() *gin.Engine {
	engine := gin.New()
	// engine := gin.Default()
	// engine.Use(gin.Logger())
	engine.Use(cfg.APILogger)
	engine.Use(gin.Recovery())

	if cfg.DebugOn {
		engine.Use(middleware.RequestDebugLogger())
	}

	// Set a lower memory limit for multipart forms (default 32M)
	engine.MaxMultipartMemory = 100 << 20 // 100MiB

	// engine := gin.Default()
	apiGroup := engine.Group("/api")
	for _, register := range registers {
		register(apiGroup)
	}

	return engine
}

func InitTemplateRouter(engine *gin.Engine) {
	// Load index html
	engine.LoadHTMLGlob(cfg.TemplateDirPattern)
	engine.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "feichai",
		})
	})

	// Serve static files
	engine.Static(cfg.StaticURL, cfg.StaticDir)
	engine.StaticFile("/favicon.ico", fmt.Sprintf("%s/favicon.ico", cfg.StaticDir))

	// Serve media files
	engine.StaticFS(cfg.MediaURL, http.Dir(cfg.MediaDir))

}
