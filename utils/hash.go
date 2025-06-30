package utils

import (
	"Library/logger"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword генерирует bcrypt-хеш из plainPassword
func HashPassword(plainPassword string) (string, error) {
	logger.Debug.Println("HashPassword: start hashing password")
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error.Printf("HashPassword: error generating hash: %v", err)
		return "", err
	}
	logger.Info.Println("HashPassword: password hashed successfully")
	return string(hashBytes), nil
}
