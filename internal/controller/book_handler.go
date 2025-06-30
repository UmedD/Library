package controller

import (
	"net/http"
	"strconv"

	"Library/internal/models"
	"Library/internal/service"
	"Library/logger"

	"github.com/gin-gonic/gin"
)

// @Summary     Список книг
// @Description Возвращает все книги с вложенным именем автора
// @Tags        books
// @Produce     json
// @Success     200 {array} models.Book
// @Failure     500 {object} models.ErrorResponse
// @Security    ApiKeyAuth
// @Router      /books [get]
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

// @Summary     Книга по ID
// @Description Возвращает книгу и имя автора по её ID
// @Tags        books
// @Produce     json
// @Param       id   path      int  true  "ID книги"
// @Success     200  {object}  models.Book
// @Failure     400 {object} models.ErrorResponse
// @Failure     404 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Security    ApiKeyAuth
// @Router      /books/{id} [get]
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

type createBookInput struct {
	Name     string `json:"name"     binding:"required"`
	Title    string `json:"title"    binding:"required"`
	AuthorID int    `json:"author_id" binding:"required"`
}

// @Summary     Создать книгу
// @Description Добавляет новую книгу (требуется роль admin)
// @Tags        books
// @Accept      json
// @Produce     json
// @Param       book  body      controller.createBookInput  true  "Поля новой книги"
// @Success     201   {object}  models.Book
// @Failure     400 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Security    ApiKeyAuth
// @Router      /books [post]
func createBook(c *gin.Context) {
	var in createBookInput
	if err := c.BindJSON(&in); err != nil {
		logger.Error.Printf("createBook: bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := models.Book{
		Name:     in.Name,
		Title:    in.Title,
		AuthorID: in.AuthorID,
	}

	if err := service.CreateBook(&b); err != nil {
		logger.Error.Printf("createBook: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info.Printf("createBook: created book ID=%d title=%q", b.ID, b.Title)
	c.JSON(http.StatusCreated, b)
}

// @Summary     Обновить книгу
// @Description Обновляет данные книги по ID (требуется роль admin)
// @Tags        books
// @Accept      json
// @Produce     json
// @Param       id    path      int                        true  "ID книги"

// @Success     200   {object}  models.Book
// @Failure     400 {object} models.ErrorResponse
// @Failure     404 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Security    ApiKeyAuth
// @Router      /books/{id} [put]
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

// @Summary     Удалить книгу
// @Description Удаляет книгу по ID (требуется роль admin)
// @Tags        books
// @Produce     json
// @Param       id   path      int  true  "ID книги"
// @Success     204 {string}  string  "No Content"
// @Failure     400 {object} models.ErrorResponse
// @Failure     404 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Security    ApiKeyAuth
// @Router      /books/{id} [delete]
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

// @Summary     Поиск книг по названию
// @Description Возвращает книги, в названии которых встречается фрагмент
// @Tags        books
// @Produce     json
// @Param       name  query     string  true  "Фрагмент в названии"
// @Success     200   {array}   models.Book
// @Failure     400 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Security    ApiKeyAuth
// @Router      /books/search [get]
func searchBooksByName(c *gin.Context) {
	fragment := c.Query("name")
	if fragment == "" {
		logger.Warn.Println("searchBooksByName: missing query param 'name'")
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'name' is required"})
		return
	}

	books, err := service.SearchBooksByName(fragment)
	if err != nil {
		logger.Error.Printf("searchBooksByName: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info.Printf("searchBooksByName: returned %d books for fragment=%q", len(books), fragment)
	c.JSON(http.StatusOK, books)
}
