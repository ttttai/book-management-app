package services

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
)

type IBookService interface {
	GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error)
	CreateBook(book *entities.Book) (*entities.Book, error)
	CreateBooks(books *[]entities.Book) (*[]entities.Book, error)
	CreateBookAuthors(bookAuthors *[]entities.BookAuthor) (*[]entities.BookAuthor, error)
	CreateBookSubjects(bookSubjects *[]entities.BookSubject) (*[]entities.BookSubject, error)
	GetBooksByTitle(title string) (*[]entities.Book, error)
	GetBookByISBN(isbn int) (*entities.Book, error)
	GetBookInfoByISBN(isbnSlices []int) (*[]entities.BookInfo, error)
}

type BookService struct {
	bookRepository    repositories.IBookRepository
	authorRepository  repositories.IAuthorRepository
	subjectRepository repositories.ISubjectRepository
}

func NewBookService(bookRepository repositories.IBookRepository, authorRepository repositories.IAuthorRepository, subjectRepository repositories.ISubjectRepository) IBookService {
	return &BookService{
		bookRepository:    bookRepository,
		authorRepository:  authorRepository,
		subjectRepository: subjectRepository,
	}
}

func (bs *BookService) GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error) {
	bookInfo, err := bs.bookRepository.GetBooksFromNdlApi(title, maxNum)
	if err != nil {
		return nil, err
	}

	return bookInfo, nil
}

func (bs *BookService) CreateBook(book *entities.Book) (*entities.Book, error) {
	result, err := bs.bookRepository.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return result, err
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

func (bs *BookService) CreateBookAuthors(bookAuthors *[]entities.BookAuthor) (*[]entities.BookAuthor, error) {
	if len(*bookAuthors) == 0 {
		return nil, nil
	}

	result, err := bs.bookRepository.CreateBookAuthors(bookAuthors)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (bs *BookService) CreateBookSubjects(bookSubjects *[]entities.BookSubject) (*[]entities.BookSubject, error) {
	if len(*bookSubjects) == 0 {
		return nil, nil
	}

	result, err := bs.bookRepository.CreateBookSubjects(bookSubjects)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (bs *BookService) GetBooksByTitle(title string) (*[]entities.Book, error) {
	result, err := bs.bookRepository.GetBooksByTitle(title)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bs *BookService) GetBookByISBN(isbn int) (*entities.Book, error) {
	result, err := bs.bookRepository.GetBookByISBN(isbn)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bs *BookService) GetBookInfoByISBN(isbnSlices []int) (*[]entities.BookInfo, error) {
	result, err := bs.bookRepository.GetBookInfoByISBN(isbnSlices)
	if err != nil {
		return nil, err
	}

	return result, nil
}
