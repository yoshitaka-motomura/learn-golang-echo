package database

import (
	"log"

	"github.com/yoshitaka-motomura/learn-golang-echo/config"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(config config.Config) {
	

	// MySQLデータベースに接続
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogLevel),
	})
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
