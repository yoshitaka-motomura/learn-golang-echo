package logging

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger *slog.Logger
	once   sync.Once
)

// InitLogger initializes the global logger to output to stdout
func InitLogger() {
	once.Do(func() {
		// 標準出力にJSON形式で出力するロガーを設定
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})
}

// Logger returns the global logger
func Logger() *slog.Logger {
	return logger
}
