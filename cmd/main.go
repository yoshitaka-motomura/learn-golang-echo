package main

import (
	"log/slog"
	"os"

	"github.com/yoshitaka-motomura/learn-golang-echo/config"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/database"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/logging"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/models"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/server"
)

func main() {
    cfg := config.LoadConfig()
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    logging.InitLogger(logger)

    database.Connect()

    database.MigrateDB(&models.Todo{})

    s := server.NewServer(logging.Logger(), true)

    if err := s.Start(":" + cfg.Port); err != nil {
        logging.Logger().Error("Failed to start server", "error", err)
    }
}
