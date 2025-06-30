package middleware

import (
	"net/http"
	"strings"

	"Library/logger"
	"Library/utils"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey  = "Authorization"
	ctxUserIDKey   = "userID"
	ctxUsernameKey = "username"
	ctxRoleKey     = "userRole"
)

func JWTAuthMiddleware(c *gin.Context) {
	// 1. Получаем заголовок Authorization
	authHeader := c.GetHeader(authHeaderKey)
	logger.Debug.Printf("JWTAuthMiddleware: incoming %s %s header=%q",
		c.Request.Method, c.Request.URL.Path, authHeader)

	if authHeader == "" {
		logger.Warn.Println("JWTAuthMiddleware: missing Authorization header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}

	// 2. Ожидаем формат "Bearer <token>"
	parts := strings.Fields(authHeader)
	if len(parts) != 2 || parts[0] != "Bearer" {
		logger.Warn.Printf("JWTAuthMiddleware: invalid header format %q", authHeader)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
		return
	}
	tokenString := parts[1]
	logger.Debug.Println("JWTAuthMiddleware: token extracted")

	// 3. Парсим и верифицируем токен
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		logger.Warn.Printf("JWTAuthMiddleware: token parse/validate failed: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token: " + err.Error()})
		return
	}
	logger.Info.Printf("JWTAuthMiddleware: token valid userID=%d username=%q role=%q",
		claims.UserID, claims.Username, claims.Role)

	// 4. Кладём в контекст userID и username и role
	c.Set(ctxUserIDKey, claims.UserID)
	c.Set(ctxUsernameKey, claims.Username)
	c.Set(ctxRoleKey, claims.Role)

	// 5. Продолжаем цепочку handlers
	c.Next()
}
