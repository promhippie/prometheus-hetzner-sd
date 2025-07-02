package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/command"
)

func main() {
	if env := os.Getenv("PROMETHEUS_HETZNER_ENV_FILE"); env != "" {
		_ = godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
