package main

import (
	"errors"

	"github.com/go-kit/kit/log/level"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/action"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"gopkg.in/urfave/cli.v2"
)

var (
	// ErrMissingOutputFile defines the error if output.file is empty.
	ErrMissingOutputFile = errors.New("Missing path for output.file")

	// ErrMissingHetznerUsername defines the error if hetzner.username is empty.
	ErrMissingHetznerUsername = errors.New("Missing required hetzner.username")

	// ErrMissingHetznerPassword defines the error if hetzner.password is empty.
	ErrMissingHetznerPassword = errors.New("Missing required hetzner.password")

	// ErrMissingAnyCredentials defines the error if no credentials are provided.
	ErrMissingAnyCredentials = errors.New("Missing any credentials")
)

// Server provides the sub-command to start the server.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Start integrated server",
		Flags: []cli.Flag{
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
		},
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
					"msg", ErrMissingOutputFile,
				)

				return ErrMissingOutputFile
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
						"msg", ErrMissingHetznerUsername,
					)

					return ErrMissingHetznerUsername
				}

				if credentials.Password == "" {
					level.Error(logger).Log(
						"msg", ErrMissingHetznerPassword,
					)

					return ErrMissingHetznerPassword
				}
			}

			if len(cfg.Target.Credentials) == 0 {
				level.Error(logger).Log(
					"msg", ErrMissingAnyCredentials,
				)

				return ErrMissingAnyCredentials
			}

			return action.Server(cfg, logger)
		},
	}
}
