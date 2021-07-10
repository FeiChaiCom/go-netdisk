package services

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	cfg "go-netdisk/config"
	"go-netdisk/gin-contrib/sessions/gormstore"
	"go-netdisk/middleware"
	"go-netdisk/services/demo"
	"go-netdisk/services/login"
	"go-netdisk/services/matter"
	"go-netdisk/services/monitor"
	"go-netdisk/services/permission"
	"go-netdisk/services/preference"
	"go-netdisk/services/user"
	"net/http"
	"time"
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

func InitRouter() *gin.Engine {
	engine := gin.New()
	// engine := gin.Default()
	engine.Use(gin.Logger())
	// engine.Use(cfg.APILogger)
	engine.Use(gin.Recovery())

	// Init session
	// store := cookie.NewStore([]byte("secret"))
	// store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store, _ := gormstore.NewStore(cfg.DB, gormstore.Options{
		TableName:       "gin_sessions",
		SkipCreateTable: false,
	}, []byte("secret"))

	// If you want periodic cleanup of expired sessions:
	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)

	engine.Use(sessions.Sessions("gin-session", store))
	if cfg.ENV.Debug {
		engine.Use(middleware.RequestDebugLogger)
	}

	// Set a lower memory limit for multipart forms (default 32M)
	engine.MaxMultipartMemory = 100 << 20 // 100MiB

	// engine := gin.Default()
	apiGroup := engine.Group("/api")
	for _, register := range registers {
		register(apiGroup)
	}

	// Init template and static files serve router
	initTemplateRouter(engine)

	return engine
}

func initTemplateRouter(engine *gin.Engine) {
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
	engine.StaticFS(cfg.MediaURL, http.Dir(cfg.ENV.MediaDir))
}
