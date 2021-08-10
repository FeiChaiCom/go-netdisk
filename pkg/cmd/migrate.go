package cmd

import (
	"github.com/urfave/cli/v2"
	"go-netdisk/pkg/db"
	"go-netdisk/pkg/db/initial"
	"go-netdisk/pkg/settings"
	"log"
)

var Migrate = &cli.Command{
	Name:        "migrate",
	Usage:       "Migrate init database",
	Description: `Backup create table and insert initial data.`,
	Action:      runMigrate,
	Flags: []cli.Flag{
		stringFlag("config", "", "Custom configuration file path", []string{"c"}),
		boolFlag("verbose, v", "Show process details", []string{"v"}),
	},
}

func runMigrate(c *cli.Context) error {
	log.Println("init database start")
	cfg := settings.GetCfg()
	cfg.LoadSettings(c)
	_, _ = db.InitDB()
	initial.InitData()
	log.Println("init database finished")
	return nil
}
