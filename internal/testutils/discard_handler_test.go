package testutils

import (
	"context"
	"testing"

	"log/slog"

	"github.com/stretchr/testify/assert"
)

func TestDiscardHandler_Enabled(t *testing.T) {
	handler := &DiscardHandler{}
	ctx := context.Background()

	tests := []struct {
		level    slog.Level
		expected bool
	}{
		{slog.LevelDebug, false},
		{slog.LevelInfo, false},
		{slog.LevelWarn, false},
		{slog.LevelError, false},
	}

	for _, tt := range tests {
		t.Run(tt.level.String(), func(t *testing.T) {
			result := handler.Enabled(ctx, tt.level)
			assert.Equal(t, tt.expected, result)
		})
	}
}