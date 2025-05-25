package utils

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ZUBERKHAN034/go-ecom/cmd/config"
	"github.com/ZUBERKHAN034/go-ecom/cmd/types"
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

// Validate JWT token string and validates it, returning the claims if valid
func ValidateJWTToken(tokenString string) (*types.TokenPayload, error) {
	secret := config.Env.JwtSecret
	// Check if the secret is set
	if secret == "" {
		return nil, errors.New("JWT secret is not set")
	}

	// Check if Bearer token is provided
	if tokenString == "" {
		return nil, errors.New("token is required")
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse the token and validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: " + token.Header["alg"].(string))
		}

		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Extract claims and convert them to TokenPayload
		if claims["id"] == nil || claims["name"] == nil || claims["email"] == nil || claims["address"] == nil {
			return nil, errors.New("token claims are incomplete")
		}

		tokenPayload := types.TokenPayload{
			ID:      uint(claims["id"].(float64)), // Convert float64 to uint
			Name:    claims["name"].(string),
			Email:   claims["email"].(string),
			Address: claims["address"].(string),
		}

		return &tokenPayload, nil
	}

	// If claims are not of type jwt.MapClaims, return an error
	return nil, errors.New("failed to parse token claims")
}
