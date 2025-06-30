package repository

import (
	"Library/internal/db"
	"Library/internal/models"
	"Library/logger"
)

// GetAllAuthors возвращает всех авторов.
func GetAllAuthors() ([]models.Author, error) {
	logger.Debug.Println("repo.GetAllAuthors: executing SELECT id, name FROM authors")
	var authors []models.Author
	err := db.GetDBConn().Select(&authors, `SELECT id, name FROM authors`)
	if err != nil {
		logger.Error.Printf("repo.GetAllAuthors: query error: %v", err)
		return nil, err
	}
	logger.Info.Printf("repo.GetAllAuthors: returned %d authors", len(authors))
	return authors, nil
}

// GetAuthorByID возвращает автора по ID.
func GetAuthorByID(authorID int) (models.Author, error) {
	logger.Debug.Printf("repo.GetAuthorByID: executing SELECT id, name FROM authors WHERE id=%d", authorID)
	var author models.Author
	err := db.GetDBConn().Get(&author, `SELECT id, name FROM authors WHERE id = $1`, authorID)
	if err != nil {
		logger.Error.Printf("repo.GetAuthorByID: query error id=%d: %v", authorID, err)
		return models.Author{}, err
	}
	logger.Info.Printf("repo.GetAuthorByID: found author ID=%d name=%q", author.ID, author.Name)
	return author, nil
}

// CreateAuthor добавляет нового автора.
func CreateAuthor(author *models.Author) error {
	logger.Debug.Printf("repo.CreateAuthor: executing INSERT INTO authors (name) VALUES (%q)", author.Name)
	_, err := db.GetDBConn().Exec(`INSERT INTO authors (name) VALUES ($1)`, author.Name)
	if err != nil {
		logger.Error.Printf("repo.CreateAuthor: insert error name=%q: %v", author.Name, err)
		return err
	}
	logger.Info.Printf("repo.CreateAuthor: created author name=%q", author.Name)
	return nil
}

// UpdateAuthor обновляет имя автора.
func UpdateAuthor(author *models.Author) error {
	logger.Debug.Printf("repo.UpdateAuthor: executing UPDATE authors SET name=%q WHERE id=%d", author.Name, author.ID)
	_, err := db.GetDBConn().Exec(`UPDATE authors SET name = $1 WHERE id = $2`, author.Name, author.ID)
	if err != nil {
		logger.Error.Printf("repo.UpdateAuthor: update error id=%d: %v", author.ID, err)
		return err
	}
	logger.Info.Printf("repo.UpdateAuthor: updated author ID=%d name=%q", author.ID, author.Name)
	return nil
}

// DeleteAuthorByID удаляет автора по ID.
func DeleteAuthorByID(authorID int) error {
	logger.Debug.Printf("repo.DeleteAuthorByID: executing DELETE FROM authors WHERE id=%d", authorID)
	_, err := db.GetDBConn().Exec(`DELETE FROM authors WHERE id = $1`, authorID)
	if err != nil {
		logger.Error.Printf("repo.DeleteAuthorByID: delete error id=%d: %v", authorID, err)
		return err
	}
	logger.Info.Printf("repo.DeleteAuthorByID: deleted author ID=%d", authorID)
	return nil
}

func SearchAuthorsByName(fragment string) ([]models.Author, error) {
	logger.Debug.Printf("repo.SearchAuthorsByName: executing SELECT for fragment=%q", fragment)

	query := `
        SELECT id, name
          FROM authors
         WHERE name ILIKE '%' || $1 || '%'
    `
	var authors []models.Author
	err := db.GetDBConn().Select(&authors, query, fragment)
	if err != nil {
		logger.Error.Printf("repo.SearchAuthorsByName: query error fragment=%q: %v", fragment, err)
		return nil, translateError(err)
	}

	logger.Info.Printf("repo.SearchAuthorsByName: found %d authors matching %q", len(authors), fragment)
	return authors, nil
}
