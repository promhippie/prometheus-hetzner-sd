package command

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"gopkg.in/yaml.v3"
)

var (
	// ErrConfigFormatInvalid defines the error if ext is unsupported.
	ErrConfigFormatInvalid = errors.New("config extension is not supported")
)

func setupLogger(cfg *config.Config) *slog.Logger {
	if cfg.Logs.Pretty {
		return slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: loggerLevel(cfg),
			}),
		)
	}

	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: loggerLevel(cfg),
		}),
	)
}

func loggerLevel(cfg *config.Config) slog.Leveler {
	switch strings.ToLower(cfg.Logs.Level) {
	case "error":
		return slog.LevelError
	case "warn":
		return slog.LevelWarn
	case "info":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	}

	return slog.LevelInfo
}

func readConfig(file string, cfg *config.Config) error {
	if file == "" {
		return nil
	}

	content, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	switch strings.ToLower(filepath.Ext(file)) {
	case ".yaml", ".yml":
		if err = yaml.Unmarshal(content, cfg); err != nil {
			return err
		}
	case ".json":
		if err = json.Unmarshal(content, cfg); err != nil {
			return err
		}
	default:
		return ErrConfigFormatInvalid
	}

	return nil
}
