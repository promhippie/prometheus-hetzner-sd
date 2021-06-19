package command

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log/level"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"github.com/urfave/cli/v2"
)

// Health provides the sub-command to perform a health check.
func Health(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "health",
		Usage: "Perform health checks",
		Flags: HealthFlags(cfg),
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

			resp, err := http.Get(
				fmt.Sprintf(
					"http://%s/healthz",
					cfg.Server.Addr,
				),
			)

			if err != nil {
				level.Error(logger).Log(
					"msg", "Failed to request health check",
					"err", err,
				)

				return err
			}

			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				level.Error(logger).Log(
					"msg", "Health check seems to be in bad state",
					"err", err,
					"code", resp.StatusCode,
				)

				return err
			}

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
			EnvVars:     []string{"PROMETHEUS_HETZNER_WEB_ADDRESS"},
			Destination: &cfg.Server.Addr,
		},
		&cli.StringFlag{
			Name:        "hetzner.config",
			Value:       "",
			Usage:       "Path to Hetzner configuration file",
			EnvVars:     []string{"PROMETHEUS_HETZNER_CONFIG"},
			Destination: nil,
		},
	}
}
