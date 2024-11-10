package main

import (
	"github.com/yoshitaka-motomura/learn-golang-echo/config"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/database"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/logging"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/server"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/todos/models"
)

func main() {
    cfg := config.LoadConfig()

    // ロガーの初期化（ファイル出力設定も内部で行われる）
    logging.InitLogger()

    // データベースの接続とマイグレーション
    database.Connect()
    database.MigrateDB(&models.Todo{})

    // サーバーの起動
    s := server.NewServer(logging.Logger(), true)
    if err := s.Start(":" + cfg.Port); err != nil {
        logging.Logger().Error("Failed to start server", "error", err)
    }
}
