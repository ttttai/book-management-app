package services

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IBookService interface {
	SearchBooks(title string, maxNum int) (*[]entities.Book, error)
}

type BookService struct {
	bookRepository repositories.IBookRepository
}

func NewBookService(bookRepository repositories.IBookRepository) IBookService {
	return &BookService{
		bookRepository: bookRepository,
	}
}

func (bs *BookService) SearchBooks(title string, maxNum int) (*[]entities.Book, error) {
	books, err := bs.bookRepository.GetBooksFromNdlApi(title, maxNum)
	if err != nil {
		return nil, err
	}

	return books, nil
}
