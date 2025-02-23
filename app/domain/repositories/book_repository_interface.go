package repositories

import "github.com/ttttai/golang/domain/entities"

type IBookRepository interface {
	GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error)
	// GetBooksByTitle(title string) (*[]entities.BookInfo)
	CreateBooks(books *[]entities.Book) (*[]entities.Book, error)
}
