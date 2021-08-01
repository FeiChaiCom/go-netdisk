package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/settings"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// Config contains parameters for the New function.
type Config struct {
	ConfigFile  string
	HomePath    string
	Version     string
	Commit      string
	BuildBranch string
}

// New returns a new instance of Server.
func New(cfg Config) (*Server, error) {
	s := newServer(cfg)
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}

func newServer(cfg Config) *Server {
	rootCtx, cancel := context.WithCancel(context.Background())

	return &Server{
		context:          rootCtx,
		shutdownFn:       cancel,
		shutdownFinished: make(chan struct{}),
		cfg:              settings.GetCfg(),

		configFile:  cfg.ConfigFile,
		homePath:    cfg.HomePath,
		version:     cfg.Version,
		commit:      cfg.Commit,
		buildBranch: cfg.BuildBranch,
	}
}

// Server is responsible for managing the lifecycle of services.
type Server struct {
	context          context.Context
	shutdownFn       context.CancelFunc
	log              log.Logger
	cfg              *settings.Cfg
	shutdownOnce     sync.Once
	shutdownFinished chan struct{}
	isInitialized    bool
	mtx              sync.Mutex

	configFile  string
	homePath    string
	version     string
	commit      string
	buildBranch string
	runMode     string

	// HttpServer related
	ginEngine *gin.Engine
	httpSrv   *http.Server
	db        *gorm.DB
}

// init initializes the server and its services.
func (s *Server) init() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.isInitialized {
		return nil
	}
	s.isInitialized = true

	s.loadConfiguration()

	s.initDB()

	s.ginEngine = s.newGinEngine()
	s.registerRoutes()
	s.initServerDirs()

	return nil
}

// Run initializes and starts services. This will block until all services have
// exited. To initiate shutdown, call the Shutdown method in another goroutine.
func (s *Server) Run() error {
	defer close(s.shutdownFinished)

	if err := s.init(); err != nil {
		return err
	}

	// Stop gracefully
	s.httpSrv = &http.Server{
		Addr:           fmt.Sprintf(":%d", s.cfg.Port),
		Handler:        s.ginEngine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// ListenAndServe
	listener, err := net.Listen("tcp", s.httpSrv.Addr)
	if err != nil {
		return err
	}

	log.Printf("HTTP Server Listen: %s", s.httpSrv.Addr)

	var wg sync.WaitGroup
	wg.Add(1)

	// handle http shutdown on server context done
	go func() {
		defer wg.Done()

		<-s.context.Done()
		if err := s.httpSrv.Shutdown(context.Background()); err != nil {
			log.Printf("Failed to shutdown server: %s", err)
		}
	}()


	if err := s.httpSrv.Serve(listener); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("server was shutdown gracefully")
			return nil
		}
		return err
	}

	log.Printf("before wait")
	wg.Wait()

	return nil
}

// Shutdown initiates graceful shutdown
func (s *Server) Shutdown(ctx context.Context, reason string) (err error) {
	s.shutdownOnce.Do(func() {
		log.Printf("Shutdown started %s", reason)
		// Call cancel func to stop server
		s.shutdownFn()
		// Wait for server to shut down
		select {
		case <-s.shutdownFinished:
			s.log.Printf("Finished waiting for server to shut down")
		case <-ctx.Done():
			s.log.Printf("Timed out while waiting for server to shut down")
			err = fmt.Errorf("timeout waiting for shutdown")
		}
	})

	return
}

// ExitCode returns an exit code for a given error.
func (s *Server) ExitCode(runError error) int {
	if runError != nil {
		log.Printf("Server shutdown %s", runError)
		return 1
	}
	return 0
}

// loadConfiguration loads settings and configuration from config files.
func (s *Server) loadConfiguration() {
	s.cfg.LoadSettings()
}
