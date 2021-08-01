package main

import (
	"context"
	"fmt"
	"go-netdisk/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := server.New(server.Config{})
	if err != nil {
		log.Fatalf("create server: %s\n", err)
	}

	// Wait system interrupt signal before shutdown in 5s
	go watchSystemSignals(context.Background(), s)

	if err := s.Run(); err != nil {
		log.Fatalf("run server: %s\n", err)
	}

}

func watchSystemSignals(ctx context.Context, s *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-quit:
			log.Printf("Server receive stop signal(%s):, will shutdown now\n", sig)
			// Keep 5 seconds to finish the currently handling request
			ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			if err := s.Shutdown(ctx, fmt.Sprintf("System signal: %s", sig)); err != nil {
				log.Println("Timed out waiting for server to shut down")
			}
			log.Println("Server graceful exited")
			return
		}
	}
}
