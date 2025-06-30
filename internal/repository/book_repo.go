package repository

import (
	"Library/internal/db"
	"Library/internal/models"
	"Library/logger"
)

// GetAllBooks возвращает список всех книг.
func GetAllBooks() ([]models.Book, error) {
	logger.Debug.Println("repo.GetAllBooks: executing SELECT id, name, title, author FROM books")

	const sql = `
      SELECT 
        b.id,
        b.name,
        b.title,
        b.author_id,
        a.name AS author_name
      FROM books b
      JOIN authors a ON a.id = b.author_id
    `

	var books []models.Book
	err := db.GetDBConn().Select(&books,
		sql,
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

	const sql = `
      SELECT 
        b.id, b.name, b.title, b.author_id,
        a.name AS author_name
      FROM books b
      JOIN authors a ON a.id = b.author_id
      WHERE b.id = $1
    `

	var b models.Book
	err := db.GetDBConn().Get(&b,
		sql, bookID,
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
		book.Name, book.Title, book.AuthorID)

	const sql = `
      INSERT INTO books (name, title, author_id)
      VALUES ($1, $2, $3) RETURNING id
    `

	err := db.GetDBConn().QueryRow(
		sql, book.Name, book.Title, book.AuthorID,
	).Scan(&book.ID)
	if err != nil {
		logger.Error.Printf("repo.CreateBook: insert error name=%q title=%q: %v", book.Name, book.Title, err)
		return translateError(err)
	}

	const selectSQL = `
      SELECT 
        b.id, b.name, b.title, b.author_id,
        a.name AS author_name
      FROM books b
      JOIN authors a ON a.id = b.author_id
      WHERE b.id = $1
    `

	if err := db.GetDBConn().Get(book, selectSQL, book.ID); err != nil {
		return translateError(err)
	}

	logger.Info.Printf("repo.CreateBook: created book ID=%d title=%q", book.ID, book.Title)
	return nil
}

// UpdateBook обновляет существующую книгу.
func UpdateBook(book *models.Book) error {
	logger.Debug.Printf("repo.UpdateBook: executing UPDATE books SET name=%q, title=%q, author=%q WHERE id=%d",
		book.Name, book.Title, book.AuthorID, book.ID,
	)

	const sql = `
      UPDATE books
         SET name      = $1,
             title     = $2,
             author_id = $3
       WHERE id        = $4
    `

	_, err := db.GetDBConn().Exec(
		sql, book.Name, book.Title, book.AuthorID, book.ID,
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

func SearchBooksByName(fragment string) ([]models.Book, error) {
	logger.Debug.Printf("repo.SearchBooksByTitle: executing SELECT for fragment=%q", fragment)

	const sql = `
      SELECT 
        b.id,
        b.name,
        b.title,
        b.author_id,
        a.name AS author_name
      FROM books b
      JOIN authors a ON a.id = b.author_id
      WHERE b.name ILIKE '%' || $1 || '%'
    `
	var books []models.Book
	err := db.GetDBConn().Select(&books, sql, fragment)
	if err != nil {
		logger.Error.Printf("repo.SearchBooksByTitle: query error fragment=%q: %v", fragment, err)
		return nil, translateError(err)
	}

	logger.Info.Printf("repo.SearchBooksByTitle: found %d books matching %q", len(books), fragment)
	return books, nil
}
