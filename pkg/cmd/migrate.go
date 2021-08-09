package cmd

import (
	"log"

	"github.com/urfave/cli/v2"
)

var Migrate = &cli.Command{
	Name:        "migrate",
	Usage:       "Migrate init database",
	Description: `Backup create table and insert initial data.`,
	Action:      runMigrate,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path", []string{"c"}),
		boolFlag("verbose, v", "Show process details", []string{"v"}),
	},
}

func runMigrate(c *cli.Context) error {
	log.Println("TODO: init database once")
	return nil
}
