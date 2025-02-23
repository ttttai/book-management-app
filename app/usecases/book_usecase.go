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
	bookInfo, err := bu.bookService.GetBooksFromNdlApi(title, maxNum)
	if err != nil {
		return nil, err
	}

	var books []entities.Book
	var authors []entities.Author
	var subjects []entities.Subject
	for _, bookInfoItem := range *bookInfo {
		books = append(books, bookInfoItem.Book)

		for _, authorItem := range bookInfoItem.Authors {
			if authorItem.Name != "" {
				authors = append(authors, authorItem)
			}
		}

		for _, subjectItem := range bookInfoItem.Subjects {
			if subjectItem.SubjectName != "" {
				subjects = append(subjects, subjectItem)
			}
		}
	}

	_, createBooksErr := bu.bookService.CreateBooks(&books)
	if createBooksErr != nil {
		return nil, createBooksErr
	}
	_, createAuthorsErr := bu.authorService.CreateAuthors(&authors)
	if createAuthorsErr != nil {
		return nil, createAuthorsErr
	}
	_, createSubjectErr := bu.subjectService.CreateSubjects(&subjects)
	if createSubjectErr != nil {
		return nil, createSubjectErr
	}

	return bookInfo, nil
}
