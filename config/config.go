package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Dsn  string
}

type DatabaseConfig struct {
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
	Host     string
}

func NewDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Port:     "3306",
		User:     "user",
		Password: "password",
		Name:     "todo_db",
		Driver:   "mysql",
		Host:     "localhost",
	}
}

func (db *DatabaseConfig) ApplyEnv() {
	envVars := []struct {
		envKey    string
		fieldName string
	}{
		{"APP_DB_PORT", "Port"},
		{"APP_DB_USER", "User"},
		{"APP_DB_PASSWORD", "Password"},
		{"APP_DB_NAME", "Name"},
		{"DATABASE_DRIVER", "Driver"},
		{"APP_DB_HOST", "Host"},
	}

	val := reflect.ValueOf(db).Elem()
	for _, e := range envVars {
		if envValue := os.Getenv(e.envKey); envValue != "" {
			field := val.FieldByName(e.fieldName)
			if field.IsValid() && field.CanSet() {
				field.SetString(envValue)
			}
		}
	}
}

func (db DatabaseConfig) DSN() string {
	switch db.Driver {
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.User, db.Password, db.Host, db.Port, db.Name,
		)
	case "postgres":
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			db.User, db.Password, db.Host, db.Port, db.Name,
		)
	default:
		log.Fatalf("Unsupported database driver: %s", db.Driver)
		return ""
	}
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from environment variables")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "1323" // デフォルトのポート番号
	}

	dbConfig := NewDatabaseConfig()
	dbConfig.ApplyEnv()

	return Config{
		Port: port,
		Dsn:  dbConfig.DSN(),
	}
}
