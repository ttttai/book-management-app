package services

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IBookService interface {
	SearchBooks(title string, maxNum int) (*[]entities.BookInfo, error)
}

type BookService struct {
	bookRepository repositories.IBookRepository
}

func NewBookService(bookRepository repositories.IBookRepository) IBookService {
	return &BookService{
		bookRepository: bookRepository,
	}
}

func (bs *BookService) SearchBooks(title string, maxNum int) (*[]entities.BookInfo, error) {
	bookInfo, err := bs.bookRepository.GetBooksFromNdlApi(title, maxNum)
	if err != nil {
		return nil, err
	}

	var books []entities.Book

	for _, bookInfoItem := range *bookInfo {
		books = append(books, bookInfoItem.Book)
	}

	_, error := bs.bookRepository.CreateBooks(&books)
	if error != nil {
		return nil, error
	}

	return bookInfo, nil
}
