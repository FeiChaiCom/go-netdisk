package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	cfg "go-netdisk/config"
	"go-netdisk/services"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Init gin log to file and stdout
func InitGin() {
	log.Println("init gin log to gin.log and stdout...")
	f, _ := os.Create(cfg.ENV.LogFile)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.Println("init file upload dir...")
	if _, err := os.Stat(cfg.ENV.MediaDir); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.ENV.MediaDir, 0755); err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(cfg.ENV.MatterRoot); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.ENV.MatterRoot, 0755); err != nil {
			panic(err)
		}
	}

	if !cfg.ENV.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Global init for gin: logger/runmode
	InitGin()

	// Init mysql connection
	if err := cfg.InitDB(); err != nil {
		panic(err)
	}

	// Init url router for apis
	router := services.InitRouter()
	// _ = router.Run(fmt.Sprintf(":%d", cfg.Port))

	// Stop gracefully
	{
		srv := &http.Server{
			Addr:           fmt.Sprintf(":%d", cfg.ENV.Port),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 1 << 20, // 1MB
		}

		// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
		go func() {
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Listen for the interrupt signal.
		<-ctx.Done()

		// Restore default behavior on the interrupt signal and notify user of shutdown.
		stop()
		log.Println("shutting down gracefully, press Ctrl+C again to force")

		// The context is used to inform the server it has 5 seconds to finish the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("Server forced to shutdown: ", err)
		}

		log.Println("Server exiting")
	}
}
