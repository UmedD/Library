package middleware

import (
	"Library/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminOnly(c *gin.Context) {
	role := c.GetString(ctxRoleKey)
	logger.Debug.Printf("AdminOnly: checking role=%q for %s %s",
		role, c.Request.Method, c.Request.URL.Path)

	if role == "" {
		logger.Warn.Println("AdminOnly: role not found in token")
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not found in token"})
		return
	}

	if role != "admin" {
		logger.Warn.Printf("AdminOnly: access denied for role=%q", role)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin privileges required"})
		return
	}

	logger.Info.Println("AdminOnly: access granted for admin")
	c.Next()
}
