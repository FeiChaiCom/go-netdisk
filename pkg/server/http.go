package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/db"
	"go-netdisk/pkg/middlewares"
	"go-netdisk/pkg/services"
	"go-netdisk/pkg/sessions/gormstore"
	"go-netdisk/pkg/settings"
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

// newSession init session for gin
func (s *Server) newSession(engine *gin.Engine) {
	// store := cookie.NewStore([]byte("secret"))
	// store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	s.store, _ = gormstore.NewStore(db.DB, gormstore.Options{
		TableName:       "gin_sessions",
		SkipCreateTable: false,
	}, []byte("secret"))

	engine.Use(sessions.Sessions("gin-session", s.store))

	// If you want periodic cleanup of expired sessions:
	go s.store.PeriodicCleanup(1*time.Second, s.shutdownFinished)
}

func (s *Server) newGin() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// Set memory limit for multipart forms (default 32M)
	engine.MaxMultipartMemory = 100 << 20 // 100MiB

	if f, err := os.Create(s.cfg.LogFile); err == nil {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	if s.cfg.Debug {
		engine.Use(settings.APILogger)
		engine.Use(middlewares.RequestDebugLogger)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	s.newSession(engine)

	return engine
}

func (s *Server) initDB() {
	s.db, _ = db.InitDB()
	// initial.InitData()
}

func (s *Server) registerRoutes() {
	services.InitRouter(s.gin)
}
