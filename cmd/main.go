package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/yoshitaka-motomura/api-learn-go/config"
	"github.com/yoshitaka-motomura/api-learn-go/internal/server"
)

func main() {
    // 設定をロード
    cfg := config.LoadConfig()

    // JSON形式のロガーを作成し、デフォルトとして設定
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    slog.SetDefault(logger)

    // サーバーを初期化し、ロガーを渡す
    s := server.NewServer(logger)

    // サーバーを起動
    log.Printf("Starting server on port %s...", cfg.Port)
    if err := s.Start(":" + cfg.Port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
