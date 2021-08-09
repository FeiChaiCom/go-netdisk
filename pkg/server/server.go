package server

import (
	"context"
	"errors"
	"fmt"
	"go-netdisk/pkg/sessions/gormstore"
	"go-netdisk/pkg/settings"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

// Config contain parameters for New func.
type Config struct {
	CliContext *cli.Context
	Version    string
	Commit     string
}

// New return a new instance of Server.
func New(cfg Config) (*Server, error) {
	s := newServer(cfg)
	if err := s.init(cfg.CliContext); err != nil {
		return nil, err
	}
	return s, nil
}

func newServer(cfg Config) *Server {
	return &Server{
		shutdownFinished: make(chan struct{}),
		cfg:              settings.GetCfg(),

		version: cfg.Version,
		commit:  cfg.Commit,
	}
}

// Server is responsible for manage httpserver.
type Server struct {
	gin   *gin.Engine
	srv   *http.Server
	db    *gorm.DB
	store gormstore.Store

	shutdownFinished chan struct{}

	cfg     *settings.Cfg
	version string
	commit  string
	runMode string
}

// init initializes the httpserver
func (s *Server) init(c *cli.Context) error {
	s.loadConfig(c)

	s.initDB()

	s.gin = s.newGin()
	s.registerRoutes()
	s.initServerDirs()

	return nil
}

// Run initializes and start httpserver, block until httpserver exited.
func (s *Server) Run() error {
	// Stop gracefully
	s.srv = &http.Server{
		Addr:           fmt.Sprintf(":%d", s.cfg.Port),
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

// loadConfig loads settings and configuration from config files.
func (s *Server) loadConfig(c *cli.Context) {
	s.cfg.LoadSettings(c)
}
