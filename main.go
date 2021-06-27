package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gaomugong/go-netdisk/apps"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Init url router for apis
	router := apps.InitAPIRouter()

	// Init template and static files serve router
	apps.InitTemplateRouter(router)

	// _ = router.Run(fmt.Sprintf(":%d", cfg.Port))

	cfg.InitDB()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
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

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

// Init gin log to file and stdout
func init() {
	log.Println("init gin log to gin.log and stdout...")
	f, _ := os.Create(cfg.LogFile)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.Println("init file upload dir...")
	if _, err := os.Stat(cfg.MediaDir); os.IsNotExist(err) {
		if err = os.Mkdir(cfg.MediaDir, 0755); err != nil {
			panic(err)
		}
	}
}
