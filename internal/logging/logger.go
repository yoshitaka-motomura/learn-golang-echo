package logging

import (
	"log/slog"
	"sync"
)

var (
    logger *slog.Logger
    once   sync.Once
)

// InitLogger initializes the global logger
func InitLogger(l *slog.Logger) {
    once.Do(func() {
        logger = l
    })
}

// Logger returns the global logger
func Logger() *slog.Logger {
    return logger
}
