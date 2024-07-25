package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Way to use godotenv variables
	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")

	return Config{
		PublicHost: os.Getenv("DB_PUBLIC_HOST"),
		Port:       os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddress:  os.Getenv("DB_ADDRESS"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

// Function to get Envs from OS ENVS
// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}

// 	return fallback
// }
