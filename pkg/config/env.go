package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	PublicHost string
	DBPort     string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	Port       string
	JwtSecret  string
}

var Env = initEnv()

func initEnv() Envs {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Way to use godotenv variables
	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")

	return Envs{
		PublicHost: os.Getenv("DB_PUBLIC_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddress:  os.Getenv("DB_ADDRESS"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
	}
}
