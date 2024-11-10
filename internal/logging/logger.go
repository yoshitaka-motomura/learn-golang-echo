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

func InitLogger(handler ...slog.Handler) {
	once.Do(func() {
		var h slog.Handler
		if len(handler) > 0 {
			h = handler[0]
		} else {
            // ディレクトリが存在しない場合は作成
            if _, err := os.Stat("logs"); os.IsNotExist(err) {
                os.Mkdir("logs", 0755)
                // .gitignoreにlogs/を追加
                ignore := []byte("*.*\n!.gitignore\n")
                os.WriteFile("logs/.gitignore", ignore, 0644)
            }
			file, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				slog.Error("Failed to open log file", "error", err)
				return
			}
			h = slog.NewJSONHandler(file, nil)
		}
		logger = slog.New(h)
	})
}

func Logger() *slog.Logger {
	return logger
}
