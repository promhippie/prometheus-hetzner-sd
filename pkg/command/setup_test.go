package command

import (
	"testing"

	"github.com/promhippie/prometheus-hetzner-sd/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestSetupLogger(t *testing.T) {
	logger := setupLogger(config.Load())
	assert.NotNil(t, logger)
}
