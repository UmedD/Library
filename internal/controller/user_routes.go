package controller

import (
	"Library/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes монтирует эндпоинты /users.
// Параметр svc больше не нужен, т.к. service — это набор функций.
func RegisterUserRoutes(r *gin.Engine) {
	// публичные руты
	r.GET("/users", getAllUsers)
	r.GET("/users/:id", getUserByID)

	// защищённые руты
	authBooks := r.Group("/users", middleware.JWTAuthMiddleware, middleware.AdminOnly)
	{
		authBooks.POST("", createUser)
		authBooks.PUT("/:id", updateUser)
		authBooks.DELETE("/:id", deleteUser)
	}

}
