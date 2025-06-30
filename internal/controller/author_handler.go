package controller

import (
	"net/http"
	"strconv"

	"Library/internal/models"
	"Library/internal/service"
	"Library/logger"

	"github.com/gin-gonic/gin"
)

// getAllAuthors возвращает список всех авторов.
func getAllAuthors(c *gin.Context) {
	authors, err := service.GetAllAuthors()
	if err != nil {
		logger.Error.Printf("getAllAuthors: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("getAllAuthors: returned %d authors", len(authors))
	c.JSON(http.StatusOK, authors)
}

// getAuthorByID возвращает автора по его ID.
func getAuthorByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("getAuthorByID: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	author, err := service.GetAuthorByID(id)
	if err != nil {
		logger.Error.Printf("getAuthorByID: service error for ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("getAuthorByID: returned author ID=%d name=%q", author.ID, author.Name)
	c.JSON(http.StatusOK, author)
}

// createAuthor создаёт нового автора.
func createAuthor(c *gin.Context) {
	var a models.Author
	if err := c.BindJSON(&a); err != nil {
		logger.Error.Printf("createAuthor: bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateAuthor(&a); err != nil {
		logger.Error.Printf("createAuthor: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("createAuthor: created author ID=%d name=%q", a.ID, a.Name)
	c.JSON(http.StatusCreated, a)
}

// updateAuthor обновляет данные существующего автора.
func updateAuthor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("updateAuthor: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	var a models.Author
	if err := c.BindJSON(&a); err != nil {
		logger.Error.Printf("updateAuthor: bind error for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	a.ID = id

	if err := service.UpdateAuthor(&a); err != nil {
		logger.Error.Printf("updateAuthor: service error for ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("updateAuthor: updated author ID=%d name=%q", a.ID, a.Name)
	c.JSON(http.StatusOK, a)
}

// deleteAuthor удаляет автора по ID.
func deleteAuthor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("deleteAuthor: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	if err := service.DeleteAuthorByID(id); err != nil {
		logger.Error.Printf("deleteAuthor: service error for ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("deleteAuthor: deleted author ID=%d", id)
	c.Status(http.StatusNoContent)
}
