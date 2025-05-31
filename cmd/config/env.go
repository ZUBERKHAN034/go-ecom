package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envs struct {
	DBPublicHost string
	DBUser       string
	DBPassword   string
	DBName       string
	DBCACert     string
	Port         string
	JwtSecret    string
	BaseURL      string
}

func initEnv() envs {
	// Try to load .env file, but don't fail if it doesn't exist
	// This allows the app to work both locally (with .env) and in production (with actual env vars)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return envs{
		DBPublicHost: getEnvWithDefault("DB_PUBLIC_HOST", ""),
		DBUser:       getEnvWithDefault("DB_USER", ""),
		DBPassword:   getEnvWithDefault("DB_PASSWORD", ""),
		DBName:       getEnvWithDefault("DB_NAME", ""),
		DBCACert:     getEnvWithDefault("DB_CA_CERT", ""),
		Port:         getEnvWithDefault("PORT", "8080"),
		JwtSecret:    getEnvWithDefault("JWT_SECRET", "qwertyuiopasdfghjklzxcvbnm123456"),
		BaseURL:      getEnvWithDefault("BASE_URL", "http://localhost:8080"),
	}
}

// Helper function to get environment variable with default value
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var Env = initEnv()
