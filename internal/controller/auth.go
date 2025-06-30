package controller

import (
	"Library/internal/models"
	"Library/internal/service"
	"Library/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func SignUp(c *gin.Context) {
	var in signUpInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Собираем модель без роли — её подхватит DEFAULT в БД
	user := models.User{
		Username: in.Username,
		Email:    in.Email,
		Password: in.Password, // дальше bcrypt-в сервисе
		// Role оставляем пустым
	}

	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignIn — POST /auth/sign-in.
// Принимает JSON { "username": "...", "password": "..." },
// проверяет через service.AuthenticateUser, затем генерирует JWT
// с userID, username и role из БД и возвращает его.
func SignIn(c *gin.Context) {
	var in signInInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Аутентифицируем: внутри сервиса сравнение bcrypt и чтение role
	user, err := service.AuthenticateUser(in.Username, in.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Генерируем JWT, прокидывая role из модели user
	token, err := utils.GenerateToken(
		user.ID,
		user.Username,
		user.Role,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	// Возвращаем только токен
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
