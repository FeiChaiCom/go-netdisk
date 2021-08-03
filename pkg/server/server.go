package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/pkg/sessions/gormstore"
	"go-netdisk/pkg/settings"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// Config contain parameters for New func.
type Config struct {
	ConfigFile  string
	HomePath    string
	Version     string
	Commit      string
	BuildBranch string
}

// New return a new instance of Server.
func New(cfg Config) (*Server, error) {
	s := newServer(cfg)
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}

func newServer(cfg Config) *Server {
	return &Server{
		shutdownFinished: make(chan struct{}),
		cfg:              settings.GetCfg(),
		Wg:               &sync.WaitGroup{},

		configFile:  cfg.ConfigFile,
		homePath:    cfg.HomePath,
		version:     cfg.Version,
		commit:      cfg.Commit,
		buildBranch: cfg.BuildBranch,
	}
}

// Server is responsible for manage httpserver.
type Server struct {
	gin              *gin.Engine
	srv              *http.Server
	db               *gorm.DB
	cfg              *settings.Cfg
	store            gormstore.Store
	shutdownFinished chan struct{}
	isInitialized    bool
	mtx              sync.Mutex
	Wg               *sync.WaitGroup

	configFile  string
	homePath    string
	version     string
	commit      string
	buildBranch string
	runMode     string
}

// init initializes the httpserver
func (s *Server) init() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.isInitialized {
		return nil
	}
	s.isInitialized = true

	s.loadConfiguration()

	s.initDB()

	s.gin = s.newGin()
	s.registerRoutes()
	s.initServerDirs()

	return nil
}

// Run initializes and start httpserver, block until httpserver exited.
func (s *Server) Run() error {
	if err := s.init(); err != nil {
		return err
	}

	// Stop gracefully
	s.srv = &http.Server{
		Addr:           fmt.Sprintf("localhost:%d", s.cfg.Port),
		Handler:        s.gin,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// ListenAndServe
	listener, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}

	log.Printf("Server listen at: %s", listener.Addr().String())
	if err := s.srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server serve error: %s\n", err)
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context, reason string) error {
	defer close(s.shutdownFinished)

	log.Printf("Shutdown started, reason: %s\n", reason)
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Printf("Failed to shutdown server: %s\n", err)
		return err
	}

	// s.shutdownFinished <- struct{}{}

	return nil
}

// loadConfiguration loads settings and configuration from config files.
func (s *Server) loadConfiguration() {
	s.cfg.LoadSettings()
}
