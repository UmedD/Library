package utils

import (
	"Library/internal/config"
	"Library/logger"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"time"
)

// CustomClaims — поля, которые храним в токене.
type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken генерирует JWT с настройками из config.AppSettings.AuthParams.
func GenerateToken(userID int, username, role string) (string, error) {
	logger.Debug.Printf("GenerateToken: start for userID=%d username=%q role=%q", userID, username, role)

	// Достаём параметры из конфига
	auth := config.AppSettings.AuthParams

	stdClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().
			Add(time.Duration(auth.JwtTtlMinutes) * time.Minute).
			Unix(),
		Issuer: config.AppSettings.AppParams.ServerName,
	}

	// Собираем свои claims
	claims := CustomClaims{
		UserID:         userID,
		Username:       username,
		Role:           role,
		StandardClaims: stdClaims,
	}

	// Создаём и подписываем
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(auth.JwtSecretKey))
	if err != nil {
		logger.Error.Printf("GenerateToken: error signing token for userID=%d: %v", userID, err)
		return "", err
	}

	logger.Info.Printf("GenerateToken: token generated successfully for userID=%d", userID)
	return signedToken, nil
}

// ParseToken парсит и валидирует JWT, возвращает CustomClaims.
func ParseToken(tokenString string) (*CustomClaims, error) {
	logger.Debug.Println("ParseToken: start token parsing")

	auth := config.AppSettings.AuthParams

	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Warn.Printf("ParseToken: unexpected signing method: %v", token.Header["alg"])
				return nil, errors.New("unexpected signing method")
			}
			return []byte(auth.JwtSecretKey), nil
		},
	)
	if err != nil {
		logger.Error.Printf("ParseToken: parse error: %v", err)
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		logger.Warn.Println("ParseToken: invalid token or claims type mismatch")
		return nil, errors.New("invalid token")
	}

	logger.Info.Printf("ParseToken: token validated successfully for userID=%d", claims.UserID)
	return claims, nil
}
