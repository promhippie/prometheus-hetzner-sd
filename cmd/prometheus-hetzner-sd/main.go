package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	cfg := config.Load()

	if env := os.Getenv("PROMETHEUS_HETZNER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:    "prometheus-hetzner-sd",
		Version: version.Version,
		Usage:   "Prometheus Hetzner SD",
		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log.level",
				Value:       "info",
				Usage:       "Only log messages with given severity",
				EnvVars:     []string{"PROMETHEUS_HETZNER_LOG_LEVEL"},
				Destination: &cfg.Logs.Level,
			},
			&cli.BoolFlag{
				Name:        "log.pretty",
				Value:       false,
				Usage:       "Enable pretty messages for logging",
				EnvVars:     []string{"PROMETHEUS_HETZNER_LOG_PRETTY"},
				Destination: &cfg.Logs.Pretty,
			},
		},
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

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
