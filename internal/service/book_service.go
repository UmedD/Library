package service

import (
	"Library/internal/models"
	"Library/internal/repository"
	"Library/logger"
)

// GetAllBooks возвращает все книги с логированием.
func GetAllBooks() ([]models.Book, error) {
	logger.Debug.Println("service.GetAllBooks: start")
	books, err := repository.GetAllBooks()
	if err != nil {
		logger.Error.Printf("service.GetAllBooks: error fetching books: %v", err)
		return nil, err
	}
	logger.Info.Printf("service.GetAllBooks: returned %d books", len(books))
	return books, nil
}

// GetBookByID возвращает книгу по ID с логированием.
func GetBookByID(bookID int) (models.Book, error) {
	logger.Debug.Printf("service.GetBookByID: start id=%d", bookID)
	book, err := repository.GetBookByID(bookID)
	if err != nil {
		logger.Error.Printf("service.GetBookByID: error fetching book id=%d: %v", bookID, err)
		return models.Book{}, err
	}
	logger.Info.Printf("service.GetBookByID: returned book ID=%d title=%q", book.ID, book.Title)
	return book, nil
}

// CreateBook создаёт новую книгу с логированием.
func CreateBook(book *models.Book) error {
	logger.Debug.Printf("service.CreateBook: start name=%q title=%q", book.Name, book.Title)
	err := repository.CreateBook(book)
	if err != nil {
		logger.Error.Printf("service.CreateBook: error creating book name=%q: %v", book.Name, err)
		return err
	}
	logger.Info.Printf("service.CreateBook: created book ID=%d title=%q", book.ID, book.Title)
	return nil
}

// UpdateBook обновляет книгу с логированием.
func UpdateBook(book *models.Book) error {
	logger.Debug.Printf("service.UpdateBook: start ID=%d name=%q title=%q", book.ID, book.Name, book.Title)
	err := repository.UpdateBook(book)
	if err != nil {
		logger.Error.Printf("service.UpdateBook: error updating book ID=%d: %v", book.ID, err)
		return err
	}
	logger.Info.Printf("service.UpdateBook: updated book ID=%d", book.ID)
	return nil
}

// DeleteBookByID удаляет книгу по ID с логированием.
func DeleteBookByID(bookID int) error {
	logger.Debug.Printf("service.DeleteBookByID: start id=%d", bookID)
	err := repository.DeleteBookByID(bookID)
	if err != nil {
		logger.Error.Printf("service.DeleteBookByID: error deleting book id=%d: %v", bookID, err)
		return err
	}
	logger.Info.Printf("service.DeleteBookByID: deleted book ID=%d", bookID)
	return nil
}

func SearchBooksByName(fragment string) ([]models.Book, error) {
	logger.Debug.Printf("service.SearchBooksByName: start fragment=%q", fragment)
	books, err := repository.SearchBooksByName(fragment)
	if err != nil {
		logger.Error.Printf("service.SearchBooksByName: error searching books fragment=%q: %v", fragment, err)
		return nil, err
	}
	logger.Info.Printf("service.SearchBooksByName: found %d books matching %q", len(books), fragment)
	return books, nil
}
