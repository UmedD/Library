package controller

import (
	"net/http"
	"strconv"

	"Library/internal/models"
	"Library/internal/service"
	"Library/logger"

	"github.com/gin-gonic/gin"
)

// getAllBooks возвращает список всех книг.
func getAllBooks(c *gin.Context) {
	books, err := service.GetAllBooks()
	if err != nil {
		logger.Error.Printf("getAllBooks: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("getAllBooks: returned %d books", len(books))
	c.JSON(http.StatusOK, books)
}

// getBookByID возвращает книгу по её ID.
func getBookByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("getBookByID: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	book, err := service.GetBookByID(id)
	if err != nil {
		logger.Error.Printf("getBookByID: service error for ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("getBookByID: returned book ID=%d title=%q", book.ID, book.Title)
	c.JSON(http.StatusOK, book)
}

// createBook создаёт новую книгу.
func createBook(c *gin.Context) {
	var b models.Book
	if err := c.BindJSON(&b); err != nil {
		logger.Error.Printf("createBook: bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateBook(&b); err != nil {
		logger.Error.Printf("createBook: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("createBook: created book ID=%d title=%q", b.ID, b.Title)
	c.JSON(http.StatusCreated, b)
}

// updateBook обновляет данные существующей книги.
func updateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("updateBook: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	var b models.Book
	if err := c.BindJSON(&b); err != nil {
		logger.Error.Printf("updateBook: bind error for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b.ID = id

	if err := service.UpdateBook(&b); err != nil {
		logger.Error.Printf("updateBook: service error for ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("updateBook: updated book ID=%d title=%q", b.ID, b.Title)
	c.JSON(http.StatusOK, b)
}

// deleteBook удаляет книгу по ID.
func deleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("deleteBook: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	if err := service.DeleteBookByID(id); err != nil {
		logger.Error.Printf("deleteBook: service error for ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("deleteBook: deleted book ID=%d", id)
	c.Status(http.StatusNoContent)
}
