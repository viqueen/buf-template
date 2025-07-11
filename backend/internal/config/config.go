package config

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Config struct {
	Environment string
	LogLevel    string
	ServerPort  string
	Database    DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

func Load() *Config {
	return &Config{
		Environment: getEnvWithDefault("ENVIRONMENT", "development"),
		LogLevel:    getEnvWithDefault("LOG_LEVEL", "info"),
		ServerPort:  getEnvWithDefault("SERVER_PORT", "8080"),
		Database: DatabaseConfig{
			Host:     getEnvWithDefault("DB_HOST", "localhost"),
			Port:     getEnvWithDefault("DB_PORT", "5432"),
			Name:     getEnvWithDefault("DB_NAME", "todo"),
			User:     getEnvWithDefault("DB_USER", "todo"),
			Password: getEnvWithDefault("DB_PASSWORD", "todo"),
			SSLMode:  getEnvWithDefault("DB_SSL_MODE", "disable"),
		},
	}
}

func (c *Config) IsDevelopment() bool {
	env := strings.ToLower(c.Environment)

	return env == "development" || env == "dev"
}

func (d *DatabaseConfig) ConnectionString() string {
	target := net.JoinHostPort(d.Host, d.Port)

	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		d.User, d.Password, target, d.Name, d.SSLMode)
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
