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

var Env = initEnv()

func initEnv() envs {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Way to use godotenv variables
	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")

	return envs{
		DBPublicHost: os.Getenv("DB_PUBLIC_HOST"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBCACert:     os.Getenv("DB_CA_CERT"),
		Port:         os.Getenv("PORT"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
		BaseURL:      os.Getenv("BASE_URL"),
	}
}
