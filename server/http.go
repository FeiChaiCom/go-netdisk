package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-netdisk/db"
	"go-netdisk/db/initial"
	"go-netdisk/gin-contrib/sessions/gormstore"
	"go-netdisk/middleware"
	"go-netdisk/services"
	"go-netdisk/settings"
	"io"
	"log"
	"os"
	"time"
)

// Init gin log to file and stdout
func (s *Server) initServerDirs() {
	log.Println("init file upload dir...")
	if _, err := os.Stat(settings.ENV.MediaDir); os.IsNotExist(err) {
		if err = os.Mkdir(settings.ENV.MediaDir, 0755); err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(settings.ENV.MatterRoot); os.IsNotExist(err) {
		if err = os.Mkdir(settings.ENV.MatterRoot, 0755); err != nil {
			panic(err)
		}
	}
}

func (s *Server) newGinEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	// engine := gin.Default()
	// engine.Use(cfg.APILogger)

	f, _ := os.Create(s.cfg.LogFile)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	if s.cfg.Debug {
		engine.Use(middleware.RequestDebugLogger)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set a lower memory limit for multipart forms (default 32M)
	engine.MaxMultipartMemory = 100 << 20 // 100MiB

	// Init session
	// store := cookie.NewStore([]byte("secret"))
	// store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store, _ := gormstore.NewStore(db.DB, gormstore.Options{
		TableName:       "gin_sessions",
		SkipCreateTable: false,
	}, []byte("secret"))

	// If you want periodic cleanup of expired sessions:
	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)
	engine.Use(sessions.Sessions("gin-session", store))

	return engine
}

func (s *Server) initDB() {
	s.db, _ = db.InitDB()
	initial.InitData()
}

func (s *Server) registerRoutes() {
	services.InitRouter(s.ginEngine)
}
