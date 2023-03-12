package command

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"gopkg.in/yaml.v3"
)

var (
	// ErrConfigFormatInvalid defines the error if ext is unsupported.
	ErrConfigFormatInvalid = errors.New("config extension is not supported")
)

func setupLogger(cfg *config.Config) log.Logger {
	var logger log.Logger

	if cfg.Logs.Pretty {
		logger = log.NewSyncLogger(
			log.NewLogfmtLogger(os.Stdout),
		)
	} else {
		logger = log.NewSyncLogger(
			log.NewJSONLogger(os.Stdout),
		)
	}

	switch strings.ToLower(cfg.Logs.Level) {
	case "error":
		logger = level.NewFilter(logger, level.AllowError())
	case "warn":
		logger = level.NewFilter(logger, level.AllowWarn())
	case "info":
		logger = level.NewFilter(logger, level.AllowInfo())
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	default:
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	return log.With(
		logger,
		"ts", log.DefaultTimestampUTC,
	)
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
