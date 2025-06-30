package controller

import (
	"net/http"
	"strconv"

	"Library/internal/models"
	"Library/internal/service"
	"Library/logger"

	"github.com/gin-gonic/gin"
)

// @Summary     Список авторов
// @Description Возвращает всех авторов
// @Tags        authors
// @Produce     json
// @Success     200 {array} models.Author
// @Failure     500 {object} models.ErrorResponse
// @Router      /authors [get]
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

// @Summary     Автор по ID
// @Description Возвращает одного автора по его ID
// @Tags        authors
// @Produce     json
// @Param       id   path      int  true  "ID автора"
// @Success     200  {object}  models.Author
// @Failure     400 {object} models.ErrorResponse
// @Failure     404 {object} models.ErrorResponse
// @Router      /authors/{id} [get]
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

// @Summary     Создать автора
// @Description Добавляет нового автора (Admin only)
// @Tags        authors
// @Security    ApiKeyAuth
// @Accept      json
// @Produce     json
// @Param       author  body      models.Author  true  "Имя нового автора"
// @Success     201     {object}  models.Author
// @Failure     400 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Router      /authors [post]
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

// @Summary     Обновить автора
// @Description Меняет имя автора по ID (Admin only)
// @Tags        authors
// @Security    ApiKeyAuth
// @Accept      json
// @Produce     json
// @Param       id      path      int            true  "ID автора"
// @Param       author  body      models.Author  true  "Новое имя автора"
// @Success     200     {object}  models.Author
// @Failure     400 {object} models.ErrorResponse
// @Failure     404 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Router      /authors/{id} [put]
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

// @Summary     Удалить автора
// @Description Удаляет автора по ID (Admin only)
// @Tags        authors
// @Security    ApiKeyAuth
// @Produce     json
// @Param       id   path      int  true  "ID автора"
// @Success     204 {string}  string  "No Content"
// @Failure     400 {object} models.ErrorResponse
// @Failure     404 {object} models.ErrorResponse
// @Failure     500 {object} models.ErrorResponse
// @Router      /authors/{id} [delete]
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

func searchAuthorsByName(c *gin.Context) {
	fragment := c.Query("name")
	if fragment == "" {
		logger.Warn.Println("searchAuthorsByName: missing query param 'name'")
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'name' is required"})
		return
	}

	authors, err := service.SearchAuthorsByName(fragment)
	if err != nil {
		logger.Error.Printf("searchAuthorsByName: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info.Printf("searchAuthorsByName: returned %d authors for fragment=%q", len(authors), fragment)
	c.JSON(http.StatusOK, authors)
}
