package config

import (
	"os"
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
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

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

var Env = initEnv()
