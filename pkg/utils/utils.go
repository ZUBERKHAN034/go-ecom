package utils

import (
	"time"

	"github.com/ZUBERKHAN034/go-ecom/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateJWT(payload map[string]interface{}) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	payload["exp"] = expirationTime
	// Set the expiration time to 24 hours from now

	secret := config.Env.JwtSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
