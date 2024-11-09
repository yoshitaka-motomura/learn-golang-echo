package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("loads config from .env file", func(t *testing.T) {
		os.Setenv("APP_PORT", "8080")
		config := LoadConfig()
		assert.Equal(t, "8080", config.Port)
	})

	t.Run("uses default port if APP_PORT is not set", func(t *testing.T) {
		os.Unsetenv("APP_PORT")
		config := LoadConfig()
		assert.Equal(t, "1323", config.Port)
	})
}