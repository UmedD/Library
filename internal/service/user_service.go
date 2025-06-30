package service

import (
	"errors"

	"Library/internal/models"
	"Library/internal/repository"
	"Library/logger"
	"Library/utils"

	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers возвращает всех пользователей с логированием.
func GetAllUsers() ([]models.User, error) {
	logger.Debug.Println("service.GetAllUsers: start")
	users, err := repository.GetallUsers()
	if err != nil {
		logger.Error.Printf("service.GetAllUsers: error fetching users: %v", err)
		return nil, err
	}
	logger.Info.Printf("service.GetAllUsers: returned %d users", len(users))
	return users, nil
}

// GetUserByID возвращает пользователя по ID с логированием.
func GetUserByID(userID int) (models.User, error) {
	logger.Debug.Printf("service.GetUserByID: start id=%d", userID)
	user, err := repository.GetUserByID(userID)
	if err != nil {
		logger.Error.Printf("service.GetUserByID: error fetching user id=%d: %v", userID, err)
		return models.User{}, err
	}
	logger.Info.Printf("service.GetUserByID: found user ID=%d username=%q", user.ID, user.Username)
	return user, nil
}

// CreateUser создаёт нового пользователя с логированием и хешированием пароля.
func CreateUser(user *models.User) error {
	logger.Debug.Printf("service.CreateUser: start username=%q email=%q", user.Username, user.Email)
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		logger.Error.Printf("service.CreateUser: hash error: %v", err)
		return err
	}
	user.Password = hash

	if err := repository.CreateUser(user); err != nil {
		logger.Error.Printf("service.CreateUser: repository error creating user username=%q: %v", user.Username, err)
		return err
	}
	logger.Info.Printf("service.CreateUser: created user ID=%d username=%q", user.ID, user.Username)
	return nil
}

// UpdateUser обновляет данные пользователя с логированием.
func UpdateUser(user *models.User) error {
	logger.Debug.Printf("service.UpdateUser: start ID=%d username=%q email=%q role=%q", user.ID, user.Username, user.Email, user.Role)
	if err := repository.UpdateUser(user); err != nil {
		logger.Error.Printf("service.UpdateUser: error updating user ID=%d: %v", user.ID, err)
		return err
	}
	logger.Info.Printf("service.UpdateUser: updated user ID=%d username=%q", user.ID, user.Username)
	return nil
}

// DeleteUserByID удаляет пользователя по ID с логированием.
func DeleteUserByID(userID int) error {
	logger.Debug.Printf("service.DeleteUserByID: start id=%d", userID)
	if err := repository.DeleteUserByID(userID); err != nil {
		logger.Error.Printf("service.DeleteUserByID: error deleting user id=%d: %v", userID, err)
		return err
	}
	logger.Info.Printf("service.DeleteUserByID: deleted user ID=%d", userID)
	return nil
}

// AuthenticateUser проверяет учётные данные пользователя с логированием.
func AuthenticateUser(username, plainPassword string) (*models.User, error) {
	logger.Debug.Printf("service.AuthenticateUser: start username=%q", username)
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		logger.Warn.Printf("service.AuthenticateUser: user not found username=%q", username)
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(plainPassword),
	); err != nil {
		logger.Warn.Printf("service.AuthenticateUser: password mismatch for username=%q", username)
		return nil, errors.New("invalid credentials")
	}

	logger.Info.Printf("service.AuthenticateUser: authenticated userID=%d username=%q role=%q", user.ID, user.Username, user.Role)
	return user, nil
}
