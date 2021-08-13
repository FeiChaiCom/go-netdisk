package cmd

import (
	"time"

	"github.com/urfave/cli/v2"
)

func stringFlag(name, value, usage string, aliases []string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    name,
		Aliases: aliases,
		Value:   value,
		Usage:   usage,
	}
}

func boolFlag(name, usage string, aliases []string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    name,
		Aliases: aliases,
		Usage:   usage,
	}
}

//nolint:deadcode,unused
func intFlag(name string, value int, usage string, aliases []string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:    name,
		Aliases: aliases,
		Value:   value,
		Usage:   usage,
	}
}

//nolint:deadcode,unused
func durationFlag(name string, value time.Duration, usage string, aliases []string) *cli.DurationFlag {
	return &cli.DurationFlag{
		Name:    name,
		Aliases: aliases,
		Value:   value,
		Usage:   usage,
	}
}
