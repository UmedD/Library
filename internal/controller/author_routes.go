package controller

import (
	"Library/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthorRoutes(r *gin.Engine) {
	// публичные руты
	r.GET("/authors", getAllAuthors)
	r.GET("/authors/:id", getAuthorByID)
	r.GET("/authors/search", searchAuthorsByName)

	// защищённые руты
	authBooks := r.Group("/authors", middleware.JWTAuthMiddleware, middleware.AdminOnly)
	{
		authBooks.POST("", createAuthor)
		authBooks.PUT("/:id", updateAuthor)
		authBooks.DELETE("/:id", deleteAuthor)
	}

}
