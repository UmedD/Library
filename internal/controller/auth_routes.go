package controller

import (
	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes монтирует публичные эндпоинты /auth/sign-up и /auth/sign-in.
func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}
}
