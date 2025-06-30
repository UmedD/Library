package repository

import (
	"Library/internal/db"
	"Library/internal/models"
	"Library/logger"
)

// GetAllBooks возвращает список всех книг.
func GetAllBooks() ([]models.Book, error) {
	logger.Debug.Println("repo.GetAllBooks: executing SELECT id, name, title, author FROM books")
	var books []models.Book
	err := db.GetDBConn().Select(&books,
		`SELECT id, name, title, author FROM books`,
	)
	if err != nil {
		logger.Error.Printf("repo.GetAllBooks: query error: %v", err)
		return nil, translateError(err)
	}
	logger.Info.Printf("repo.GetAllBooks: returned %d books", len(books))
	return books, nil
}

// GetBookByID возвращает книгу по ID.
func GetBookByID(bookID int) (models.Book, error) {
	logger.Debug.Printf("repo.GetBookByID: executing SELECT id, name, title, author FROM books WHERE id=%d", bookID)
	var b models.Book
	err := db.GetDBConn().Get(&b,
		`SELECT id, name, title, author FROM books WHERE id = $1`, bookID,
	)
	if err != nil {
		logger.Error.Printf("repo.GetBookByID: query error id=%d: %v", bookID, err)
		return models.Book{}, translateError(err)
	}
	logger.Info.Printf("repo.GetBookByID: found book ID=%d title=%q", b.ID, b.Title)
	return b, nil
}

// CreateBook сохраняет новую книгу.
func CreateBook(book *models.Book) error {
	logger.Debug.Printf("repo.CreateBook: executing INSERT INTO books (name, title, author) VALUES (%q, %q, %q)",
		book.Name, book.Title, book.Author,
	)
	err := db.GetDBConn().QueryRow(
		`INSERT INTO books (name, title, author) VALUES ($1, $2, $3) RETURNING id`,
		book.Name, book.Title, book.Author,
	).Scan(&book.ID)
	if err != nil {
		logger.Error.Printf("repo.CreateBook: insert error name=%q title=%q: %v", book.Name, book.Title, err)
		return translateError(err)
	}
	logger.Info.Printf("repo.CreateBook: created book ID=%d title=%q", book.ID, book.Title)
	return nil
}

// UpdateBook обновляет существующую книгу.
func UpdateBook(book *models.Book) error {
	logger.Debug.Printf("repo.UpdateBook: executing UPDATE books SET name=%q, title=%q, author=%q WHERE id=%d",
		book.Name, book.Title, book.Author, book.ID,
	)
	_, err := db.GetDBConn().Exec(
		`UPDATE books SET name = $1, title = $2, author = $3 WHERE id = $4`,
		book.Name, book.Title, book.Author, book.ID,
	)
	if err != nil {
		logger.Error.Printf("repo.UpdateBook: exec error ID=%d: %v", book.ID, err)
		return translateError(err)
	}
	logger.Info.Printf("repo.UpdateBook: updated book ID=%d title=%q", book.ID, book.Title)
	return nil
}

// DeleteBookByID удаляет книгу по ID.
func DeleteBookByID(bookID int) error {
	logger.Debug.Printf("repo.DeleteBookByID: executing DELETE FROM books WHERE id=%d", bookID)
	_, err := db.GetDBConn().Exec(
		`DELETE FROM books WHERE id = $1`, bookID,
	)
	if err != nil {
		logger.Error.Printf("repo.DeleteBookByID: delete error ID=%d: %v", bookID, err)
		return translateError(err)
	}
	logger.Info.Printf("repo.DeleteBookByID: deleted book ID=%d", bookID)
	return nil
}
