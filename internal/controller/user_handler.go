package controller

import (
	"net/http"
	"strconv"

	"Library/internal/models"
	"Library/internal/service"
	"Library/logger"

	"github.com/gin-gonic/gin"
)

// getAllUsers отдаёт всех пользователей.
func getAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		logger.Error.Printf("getAllUsers: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("getAllUsers: returned %d users", len(users))
	c.JSON(http.StatusOK, users)
}

// getUserByID отдаёт одного пользователя по ID.
func getUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("getUserByID: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		logger.Error.Printf("getUserByID: service error for ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("getUserByID: returned user ID=%d username=%q", user.ID, user.Username)
	c.JSON(http.StatusOK, user)
}

// createUser создаёт нового пользователя.
func createUser(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		logger.Error.Printf("createUser: bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateUser(&u); err != nil {
		logger.Error.Printf("createUser: service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("createUser: created user ID=%d username=%q", u.ID, u.Username)
	c.JSON(http.StatusCreated, u)
}

// updateUser обновляет существующего пользователя.
func updateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("updateUser: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var u models.User
	if err := c.BindJSON(&u); err != nil {
		logger.Error.Printf("updateUser: bind error for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.ID = id

	if err := service.UpdateUser(&u); err != nil {
		logger.Error.Printf("updateUser: service error for ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("updateUser: updated user ID=%d username=%q", u.ID, u.Username)
	c.JSON(http.StatusOK, u)
}

// deleteUser удаляет пользователя.
func deleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Error.Printf("deleteUser: invalid ID param %q: %v", idParam, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := service.DeleteUserByID(id); err != nil {
		logger.Error.Printf("deleteUser: service error for ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info.Printf("deleteUser: deleted user ID=%d", id)
	c.Status(http.StatusNoContent)
}
