package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IBookUsecase interface {
	GetBooks(title string) (*[]entities.Book, error)
}

type BookUsecase struct {
	bookRepository repositories.IBookRepository
}

func NewBookUsecase(bookRepository repositories.IBookRepository) IBookUsecase {
	return &BookUsecase{
		bookRepository: bookRepository,
	}
}

func (bu *BookUsecase) GetBooks(title string) (*[]entities.Book, error) {
	books, err := bu.bookRepository.GetBooksFromNdlApi(title)
	if err != nil {
		return nil, err
	}

	return books, nil
}
