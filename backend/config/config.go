package config

import (
	"os"
)

type Config struct {
	HeliusAPIKey string
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBPort       string
	ServerPort   string
}

func LoadConfig() *Config {
	return &Config{
		HeliusAPIKey: getEnvOrDefault("HELIUS_API_KEY", ""),
		DBHost:       getEnvOrDefault("DB_HOST", "localhost"),
		DBUser:       getEnvOrDefault("DB_USER", "bounty"),
		DBPassword:   getEnvOrDefault("DB_PASSWORD", "bounty123"),
		DBName:       getEnvOrDefault("DB_NAME", "bounty_db"),
		DBPort:       getEnvOrDefault("DB_PORT", "5432"),
		ServerPort:   getEnvOrDefault("SERVER_PORT", "3000"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDSN() string {
	return "host=" + c.DBHost + " user=" + c.DBUser + " password=" + c.DBPassword +
		" dbname=" + c.DBName + " port=" + c.DBPort + " sslmode=disable"
}