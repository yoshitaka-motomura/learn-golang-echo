package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/yoshitaka-motomura/learn-golang-echo/config"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/server"
)

func main() {
    cfg := config.LoadConfig()
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    slog.SetDefault(logger)

    s := server.NewServer(logger, true)

    log.Printf("Starting server on port %s...", cfg.Port)
    if err := s.Start(":" + cfg.Port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
