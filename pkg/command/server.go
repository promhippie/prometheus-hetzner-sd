package command

import (
	"errors"

	"github.com/go-kit/kit/log/level"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/action"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"github.com/urfave/cli/v2"
)

// Server provides the sub-command to start the server.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Start integrated server",
		Flags: ServerFlags(cfg),
		Action: func(c *cli.Context) error {
			logger := setupLogger(cfg)

			if c.IsSet("hetzner.config") {
				if err := readConfig(c.String("hetzner.config"), cfg); err != nil {
					level.Error(logger).Log(
						"msg", "Failed to read config",
						"err", err,
					)

					return err
				}
			}

			if cfg.Target.File == "" {
				level.Error(logger).Log(
					"msg", "Missing path for output.file",
				)

				return errors.New("missing path for output.file")
			}

			if c.IsSet("hetzner.username") && c.IsSet("hetzner.password") {
				credentials := config.Credential{
					Project:  "default",
					Username: c.String("hetzner.username"),
					Password: c.String("hetzner.password"),
				}

				cfg.Target.Credentials = append(
					cfg.Target.Credentials,
					credentials,
				)

				if credentials.Username == "" {
					level.Error(logger).Log(
						"msg", "Missing required hetzner.username",
					)

					return errors.New("missing required hetzner.username")
				}

				if credentials.Password == "" {
					level.Error(logger).Log(
						"msg", "Missing required hetzner.password",
					)

					return errors.New("missing required hetzner.password")
				}
			}

			if len(cfg.Target.Credentials) == 0 {
				level.Error(logger).Log(
					"msg", "Missing any credentials",
				)

				return errors.New("missing any credentials")
			}

			return action.Server(cfg, logger)
		},
	}
}

// ServerFlags defines the available server flags.
func ServerFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "web.address",
			Value:       "0.0.0.0:9000",
			Usage:       "Address to bind the metrics server",
			EnvVars:     []string{"PROMETHEUS_HETZNER_WEB_ADDRESS"},
			Destination: &cfg.Server.Addr,
		},
		&cli.StringFlag{
			Name:        "web.path",
			Value:       "/metrics",
			Usage:       "Path to bind the metrics server",
			EnvVars:     []string{"PROMETHEUS_HETZNER_WEB_PATH"},
			Destination: &cfg.Server.Path,
		},
		&cli.StringFlag{
			Name:        "output.file",
			Value:       "/etc/prometheus/hetzner.json",
			Usage:       "Path to write the file_sd config",
			EnvVars:     []string{"PROMETHEUS_HETZNER_OUTPUT_FILE"},
			Destination: &cfg.Target.File,
		},
		&cli.IntFlag{
			Name:        "output.refresh",
			Value:       30,
			Usage:       "Discovery refresh interval in seconds",
			EnvVars:     []string{"PROMETHEUS_HETZNER_OUTPUT_REFRESH"},
			Destination: &cfg.Target.Refresh,
		},
		&cli.StringFlag{
			Name:    "hetzner.username",
			Value:   "",
			Usage:   "Username for the Hetzner API",
			EnvVars: []string{"PROMETHEUS_HETZNER_USERNAME"},
		},
		&cli.StringFlag{
			Name:    "hetzner.password",
			Value:   "",
			Usage:   "Password for the Hetzner API",
			EnvVars: []string{"PROMETHEUS_HETZNER_PASSWORD"},
		},
		&cli.StringFlag{
			Name:    "hetzner.config",
			Value:   "",
			Usage:   "Path to Hetzner configuration file",
			EnvVars: []string{"PROMETHEUS_HETZNER_CONFIG"},
		},
	}
}
