package usecases

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/services"
)

type IBookUsecase interface {
	SearchBooks(title string, maxNum int) (*[]entities.BookInfo, error)
}

type BookUsecase struct {
	bookService    services.IBookService
	authorService  services.IAuthorService
	subjectService services.ISubjectService
}

func NewBookUsecase(bookService services.IBookService, authorService services.IAuthorService, subjectService services.ISubjectService) IBookUsecase {
	return &BookUsecase{
		bookService:    bookService,
		authorService:  authorService,
		subjectService: subjectService,
	}
}

func (bu *BookUsecase) SearchBooks(title string, maxNum int) (*[]entities.BookInfo, error) {
	// TODO:データベースから取得して足りない分をAPIで除外検索
	bookInfo, err := bu.bookService.GetBooksFromNdlApi(title, maxNum)
	if err != nil {
		return nil, err
	}

	for _, bookInfoItem := range *bookInfo {
		newBook, err := bu.bookService.CreateBook(&bookInfoItem.Book)
		if err != nil {
			return nil, err
		}

		bookAuthorRelations, err := bu.authorService.GetBookAuthorRelations(newBook, &bookInfoItem.Authors)
		if err != nil {
			return nil, err
		}
		_, err = bu.bookService.CreateBookAuthors(bookAuthorRelations)
		if err != nil {
			return nil, err
		}

		bookSubjectRelations, err := bu.subjectService.GetBookSubjectRelations(newBook, &bookInfoItem.Subjects)
		if err != nil {
			return nil, err
		}
		_, err = bu.bookService.CreateBookSubjects(bookSubjectRelations)
		if err != nil {
			return nil, err
		}
	}

	return bookInfo, nil
}
