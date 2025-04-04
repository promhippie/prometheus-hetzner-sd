package command

import (
	"context"
	"os"

	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/version"
	"github.com/urfave/cli/v3"
)

// Run parses the command line arguments and executes the program.
func Run() error {
	cfg := config.Load()

	app := &cli.Command{
		Name:    "prometheus-hetzner-sd",
		Version: version.String,
		Usage:   "Prometheus Hetzner SD",
		Authors: []any{
			"Thomas Boerger <thomas@webhippie.de>",
		},
		Flags: RootFlags(cfg),
		Commands: []*cli.Command{
			Health(cfg),
			Server(cfg),
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	return app.Run(context.Background(), os.Args)
}

// RootFlags defines the available root flags.
func RootFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "log.level",
			Value:       "info",
			Usage:       "Only log messages with given severity",
			Sources:     cli.EnvVars("PROMETHEUS_HETZNER_LOG_LEVEL"),
			Destination: &cfg.Logs.Level,
		},
		&cli.BoolFlag{
			Name:        "log.pretty",
			Value:       false,
			Usage:       "Enable pretty messages for logging",
			Sources:     cli.EnvVars("PROMETHEUS_HETZNER_LOG_PRETTY"),
			Destination: &cfg.Logs.Pretty,
		},
	}
}
