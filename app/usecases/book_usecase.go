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
	var isbns []int

	// TODO:データベースから取得して足りない分をAPIで除外検索
	bookInfoFromApi, err := bu.bookService.GetBooksFromNdlApi(title, maxNum)
	if err != nil {
		return nil, err
	}

	var excludedBookInfo []entities.BookInfo
	for _, bookInfoItem := range *bookInfoFromApi {
		book, err := bu.bookService.GetBookByISBN(bookInfoItem.Book.ISBN)
		if err != nil {
			return nil, err
		}
		if book == nil {
			excludedBookInfo = append(excludedBookInfo, bookInfoItem)
		} else {
			isbns = append(isbns, book.ISBN)
		}
	}

	for _, bookInfoItem := range excludedBookInfo {
		newBook, err := bu.bookService.CreateBook(&bookInfoItem.Book)
		if err != nil {
			return nil, err
		}
		isbns = append(isbns, newBook.ISBN)

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

	bookInfo, err := bu.bookService.GetBookInfoByISBNs(isbns)

	return bookInfo, nil
}
