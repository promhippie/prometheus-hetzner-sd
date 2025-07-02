package command

import (
	"context"
	"fmt"
	"net/http"

	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"github.com/urfave/cli/v3"
)

// Health provides the sub-command to perform a health check.
func Health(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "health",
		Usage: "Perform health checks",
		Flags: HealthFlags(cfg),
		Action: func(_ context.Context, cmd *cli.Command) error {
			logger := setupLogger(cfg)

			if cmd.IsSet("hetzner.config") {
				if err := readConfig(cmd.String("hetzner.config"), cfg); err != nil {
					logger.Error("Failed to read config",
						"err", err,
					)

					return err
				}
			}

			resp, err := http.Get(
				fmt.Sprintf(
					"http://%s/healthz",
					cfg.Server.Addr,
				),
			)

			if err != nil {
				logger.Error("Failed to request health check",
					"err", err,
				)

				return err
			}

			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != 200 {
				logger.Error("Health check seems to be in bad state",
					"err", err,
					"code", resp.StatusCode,
				)

				return err
			}

			logger.Debug("Health check seems to be fine",
				"code", resp.StatusCode,
			)

			return nil
		},
	}
}

// HealthFlags defines the available health flags.
func HealthFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "web.address",
			Value:       "0.0.0.0:9000",
			Usage:       "Address to bind the metrics server",
			Sources:     cli.EnvVars("PROMETHEUS_HETZNER_WEB_ADDRESS"),
			Destination: &cfg.Server.Addr,
		},
		&cli.StringFlag{
			Name:        "hetzner.config",
			Value:       "",
			Usage:       "Path to Hetzner configuration file",
			Sources:     cli.EnvVars("PROMETHEUS_HETZNER_CONFIG"),
			Destination: nil,
		},
	}
}
