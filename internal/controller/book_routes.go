package controller

import (
	"Library/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterBookRoutes монтирует маршруты для работы с книгами.
func RegisterBookRoutes(r *gin.Engine) {
	// публичные руты
	r.GET("/books", getAllBooks)
	r.GET("/books/:id", getBookByID)
	r.GET("/books/search", searchBooksByName)

	// защищённые руты
	authBooks := r.Group("/books", middleware.JWTAuthMiddleware, middleware.AdminOnly)
	{
		authBooks.POST("", createBook)
		authBooks.PUT("/:id", updateBook)
		authBooks.DELETE("/:id", deleteBook)
	}

}
