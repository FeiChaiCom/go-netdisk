package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go-netdisk/pkg/server"
	"go-netdisk/pkg/version"

	"github.com/urfave/cli/v2"
)

var Web = &cli.Command{
	Name:        "web",
	Usage:       "Start web server",
	Description: `go-netdisk web server provide http service`,
	Action:      runWeb,
	Flags: []cli.Flag{
		intFlag("port", 5000, "Temporary port number to prevent conflict", []string{"p"}),
		stringFlag("config", "", "Custom configuration file path", []string{"c"}),
	},
}

func runWeb(c *cli.Context) error {
	s, err := server.New(server.Config{
		CliContext: c,
		Version:    version.Version,
		Commit:     version.GitHash,
	})

	if err != nil {
		log.Fatalf("Server init error: %s\n", err)
	}

	// Wait system interrupt signal before shutdown in 5s
	var wg = &sync.WaitGroup{}
	go func() {
		defer wg.Done()
		watchSystemSignals(context.Background(), s)
	}()

	if err := s.Run(); err != nil {
		log.Fatalf("Server start error: %s\n", err)
	}

	wg.Wait()
	return nil
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
