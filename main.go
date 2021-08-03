package main

import (
	"context"
	"fmt"
	"go-netdisk/pkg/server"
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
	s.Wg.Add(1)
	go func() {
		defer s.Wg.Done()
		watchSystemSignals(context.Background(), s)
	}()

	if err := s.Run(); err != nil {
		log.Fatalf("Server start error: %s\n", err)
	}

	s.Wg.Wait()
	// log.Println("exit main")
}

func watchSystemSignals(rootCtx context.Context, s *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-quit:
			log.Printf("Server receive stop signal(%s):, will shutdown now\n", sig)
			// Keep 5 seconds to finish the currently handling request
			ctx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
			defer cancel()

			if err := s.Shutdown(ctx, fmt.Sprintf("System signal: %s", sig)); err != nil {
				log.Println("Server shutdown timed out, forced stop")
				return
			}

			log.Printf("Server shutdown gracefully\n")
			return
		}
	}
}
