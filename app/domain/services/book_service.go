package services

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IBookService interface {
	GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error)
	CreateBooks(books *[]entities.Book) (*[]entities.Book, error)
}

type BookService struct {
	bookRepository repositories.IBookRepository
}

func NewBookService(bookRepository repositories.IBookRepository) IBookService {
	return &BookService{
		bookRepository: bookRepository,
	}
}

func (bs *BookService) GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error) {
	bookInfo, err := bs.bookRepository.GetBooksFromNdlApi(title, maxNum)
	if err != nil {
		return nil, err
	}

	return bookInfo, nil
}

func (bs *BookService) CreateBooks(books *[]entities.Book) (*[]entities.Book, error) {
	if len(*books) == 0 {
		return nil, nil
	}

	result, err := bs.bookRepository.CreateBooks(books)
	if err != nil {
		return nil, err
	}

	return result, err
}

// func (bs *BookService) GetBooksByTitle(title string) (*[]entities.Book, error) {

// }
