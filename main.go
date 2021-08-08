// +build go1.16

package main

import (
	"go-netdisk/pkg/cmd"
	"go-netdisk/pkg/version"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-netdisk"
	app.Usage = "A simple net-disk service"
	app.Version = version.Version
	app.Commands = []cli.Command{
		cmd.Web,
		// cmd.Backup,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
