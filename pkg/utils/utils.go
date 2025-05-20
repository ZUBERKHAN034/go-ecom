package utils

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ZUBERKHAN034/go-ecom/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a password string and returns a hashed password string
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePassword takes a hashed password and a password string and returns a boolean
func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// GenerateJWT takes a payload and returns a JWT token string
func GenerateJWT(payload any) (string, error) {
	var jwtPayload map[string]any

	// If payload is already a map, use it directly
	if p, ok := payload.(map[string]any); ok {
		jwtPayload = p
	} else {
		// Convert struct to map
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return "", errors.New("failed to serialize payload")
		}
		json.Unmarshal(payloadBytes, &jwtPayload)
	}

	// Set the expiration time to 24 hours from now
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	jwtPayload["exp"] = expirationTime

	secret := config.Env.JwtSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(jwtPayload))
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
