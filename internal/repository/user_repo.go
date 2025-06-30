package repository

import (
	"Library/internal/db"
	"Library/internal/models"
	"Library/logger"
)

// GetallUsers возвращает всех пользователей.
func GetallUsers() ([]models.User, error) {
	logger.Debug.Println("repo.GetallUsers: executing SELECT id, username, email, role FROM users")

	var users []models.User
	err := db.GetDBConn().Select(&users,
		`SELECT id, username, email, role FROM users`,
	)
	if err != nil {
		logger.Error.Printf("repo.GetallUsers: query error: %v", err)
		return nil, translateError(err)
	}
	logger.Info.Printf("repo.GetallUsers: returned %d users", len(users))
	return users, nil
}

// GetUserByID возвращает пользователя по ID.
func GetUserByID(userID int) (models.User, error) {
	logger.Debug.Printf("repo.GetUserByID: executing SELECT id, username, email, role FROM users WHERE id=%d", userID)

	var user models.User
	err := db.GetDBConn().Get(&user,
		`SELECT id, username, email, role FROM users WHERE id = $1`, userID,
	)
	if err != nil {
		logger.Error.Printf("repo.GetUserByID: query error id=%d: %v", userID, err)
		return models.User{}, translateError(err)
	}
	logger.Info.Printf("repo.GetUserByID: found user ID=%d username=%q", user.ID, user.Username)
	return user, nil
}

// CreateUser сохраняет нового пользователя.
func CreateUser(user *models.User) error {
	logger.Debug.Printf(
		"repo.CreateUser: executing INSERT INTO users (username, email, password) VALUES (%q, %q, ****)",
		user.Username, user.Email,
	)
	err := db.GetDBConn().QueryRow(
		`INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`,
		user.Username, user.Email, user.Password,
	).Scan(&user.ID)
	if err != nil {
		logger.Error.Printf("repo.CreateUser: insert error username=%q: %v", user.Username, err)
		return translateError(err)
	}
	logger.Info.Printf("repo.CreateUser: created user ID=%d username=%q", user.ID, user.Username)
	return nil
}

// UpdateUser обновляет данные пользователя.
func UpdateUser(user *models.User) error {
	logger.Debug.Printf(
		"repo.UpdateUser: executing UPDATE users SET username=%q, email=%q, role=%q WHERE id=%d",
		user.Username, user.Email, user.Role, user.ID,
	)
	_, err := db.GetDBConn().Exec(
		`UPDATE users SET username = $1, email = $2, role = $3 WHERE id = $4`,
		user.Username, user.Email, user.Role, user.ID,
	)
	if err != nil {
		logger.Error.Printf("repo.UpdateUser: exec error id=%d: %v", user.ID, err)
		return translateError(err)
	}
	logger.Info.Printf("repo.UpdateUser: updated user ID=%d username=%q", user.ID, user.Username)
	return nil
}

// DeleteUserByID удаляет пользователя по ID.
func DeleteUserByID(userID int) error {
	logger.Debug.Printf("repo.DeleteUserByID: executing DELETE FROM users WHERE id=%d", userID)
	_, err := db.GetDBConn().Exec(
		`DELETE FROM users WHERE id = $1`, userID,
	)
	if err != nil {
		logger.Error.Printf("repo.DeleteUserByID: delete error id=%d: %v", userID, err)
		return translateError(err)
	}
	logger.Info.Printf("repo.DeleteUserByID: deleted user ID=%d", userID)
	return nil
}

// GetUserByUsername возвращает пользователя по username (для аутентификации).
func GetUserByUsername(username string) (*models.User, error) {
	logger.Debug.Printf("repo.GetUserByUsername: executing SELECT id, username, email, password, role FROM users WHERE username=%q", username)

	var u models.User
	err := db.GetDBConn().Get(&u,
		`SELECT id, username, email, password, role
           FROM users
          WHERE username = $1`, username,
	)
	if err != nil {
		logger.Error.Printf("repo.GetUserByUsername: query error username=%q: %v", username, err)
		return nil, translateError(err)
	}
	logger.Info.Printf("repo.GetUserByUsername: found user ID=%d username=%q", u.ID, u.Username)
	return &u, nil
}
