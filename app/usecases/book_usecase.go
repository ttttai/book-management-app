package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/services"
)

type IBookUsecase interface {
	SearchBooks(title string, maxNum int) (*[]entities.BookInfo, error)
}

type BookUsecase struct {
	bookService services.IBookService
}

func NewBookUsecase(bookService services.IBookService) IBookUsecase {
	return &BookUsecase{
		bookService: bookService,
	}
}

func (bu *BookUsecase) SearchBooks(title string, maxNum int) (*[]entities.BookInfo, error) {
	bookInfo, err := bu.bookService.SearchBooks(title, maxNum)
	if err != nil {
		return nil, err
	}

	return bookInfo, nil
}
