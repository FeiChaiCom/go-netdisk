package apps

import (
	"fmt"
	"github.com/gaomugong/go-netdisk/apps/demo"
	"github.com/gaomugong/go-netdisk/apps/matter"
	"github.com/gaomugong/go-netdisk/apps/monitor"
	"github.com/gaomugong/go-netdisk/apps/user"
	"github.com/gaomugong/go-netdisk/common"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Register func(rg *gin.RouterGroup)

var registers = []Register{
	user.RegisterUserGroup,
	matter.RegisterMatterGroup,
	monitor.RegisterMonitorGroup,
	demo.RegisterTestGroup,
}

func InitAPIRouter() *gin.Engine {
	// engine := gin.Default()
	engine := gin.New()
	// engine.Use(gin.Logger())
	engine.Use(cfg.ApiLogger)
	engine.Use(gin.Recovery())
	engine.Use(common.LoginRequiredMiddleware())

	// Set a lower memory limit for multipart forms (default 32M)
	engine.MaxMultipartMemory = 8 << 20 // 8MiB

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
