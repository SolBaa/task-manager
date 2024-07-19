package config

import (
	"os"
	"strconv"

	"github.com/SolBaa/task-manager/internal/db"
)

type Config struct {
	ServerPort string
	DBConfig   db.Config
	JWTSecret  string
}

func LoadConfig() *Config {
	port, err := strconv.Atoi(getEnv("DB_PORT", "3306"))
	if err != nil {
		port = 3306
	}

	return &Config{
		ServerPort: getEnv("PORT", ":8080"),
		DBConfig: db.Config{
			User:     getEnv("DB_USER", "sol"),
			Password: getEnv("DB_PASSWORD", "password"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     port,
			DBName:   getEnv("DB_NAME", "task_manager"),
		},
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
