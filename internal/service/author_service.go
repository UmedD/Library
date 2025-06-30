package service

import (
	"Library/internal/models"
	"Library/internal/repository"
	"Library/logger"
)

// GetAllAuthors возвращает всех авторов с логированием.
func GetAllAuthors() ([]models.Author, error) {
	logger.Debug.Println("service.GetAllAuthors: start")
	authors, err := repository.GetAllAuthors()
	if err != nil {
		logger.Error.Printf("service.GetAllAuthors: error fetching authors: %v", err)
		return nil, err
	}
	logger.Info.Printf("service.GetAllAuthors: returned %d authors", len(authors))
	return authors, nil
}

// GetAuthorByID возвращает автора по ID с логированием.
func GetAuthorByID(authorID int) (models.Author, error) {
	logger.Debug.Printf("service.GetAuthorByID: start id=%d", authorID)
	author, err := repository.GetAuthorByID(authorID)
	if err != nil {
		logger.Error.Printf("service.GetAuthorByID: error fetching author id=%d: %v", authorID, err)
		return models.Author{}, err
	}
	logger.Info.Printf("service.GetAuthorByID: returned author ID=%d name=%q", author.ID, author.Name)
	return author, nil
}

// CreateAuthor создаёт нового автора с логированием.
func CreateAuthor(author *models.Author) error {
	logger.Debug.Printf("service.CreateAuthor: start name=%q", author.Name)
	err := repository.CreateAuthor(author)
	if err != nil {
		logger.Error.Printf("service.CreateAuthor: error creating author name=%q: %v", author.Name, err)
		return err
	}
	logger.Info.Printf("service.CreateAuthor: created author ID=%d name=%q", author.ID, author.Name)
	return nil
}

// UpdateAuthor обновляет автора с логированием.
func UpdateAuthor(author *models.Author) error {
	logger.Debug.Printf("service.UpdateAuthor: start ID=%d name=%q", author.ID, author.Name)
	err := repository.UpdateAuthor(author)
	if err != nil {
		logger.Error.Printf("service.UpdateAuthor: error updating author ID=%d: %v", author.ID, err)
		return err
	}
	logger.Info.Printf("service.UpdateAuthor: updated author ID=%d name=%q", author.ID, author.Name)
	return nil
}

// DeleteAuthorByID удаляет автора по ID с логированием.
func DeleteAuthorByID(authorID int) error {
	logger.Debug.Printf("service.DeleteAuthorByID: start id=%d", authorID)
	err := repository.DeleteAuthorByID(authorID)
	if err != nil {
		logger.Error.Printf("service.DeleteAuthorByID: error deleting author id=%d: %v", authorID, err)
		return err
	}
	logger.Info.Printf("service.DeleteAuthorByID: deleted author ID=%d", authorID)
	return nil
}

func SearchAuthorsByName(fragment string) ([]models.Author, error) {
	logger.Debug.Printf("service.SearchAuthorsByName: start fragment=%q", fragment)
	authors, err := repository.SearchAuthorsByName(fragment)
	if err != nil {
		logger.Error.Printf("service.SearchAuthorsByName: error searching authors fragment=%q: %v", fragment, err)
		return nil, err
	}
	logger.Info.Printf("service.SearchAuthorsByName: returned %d authors matching %q", len(authors), fragment)
	return authors, nil
}
