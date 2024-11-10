package database

import (
	"log"

	"github.com/yoshitaka-motomura/learn-golang-echo/config"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := config.LoadConfig()
	// 環境変数からデータベース接続情報を取得
	dsn := cfg.Dsn

	// MySQLデータベースに接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	logging.Logger().Info("Database connected successfully.")
}

// MigrateDBはデータベースのマイグレーションを行う
func MigrateDB(models ...interface{}) {
	if err := DB.AutoMigrate(models...); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
