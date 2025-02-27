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
	GetBooksByFuzzyTitle(title string) (*[]entities.Book, error)
	GetBookByISBN(isbn int) (*entities.Book, error)
	GetBookInfoByISBNs(isbnSlices []int) (*[]entities.BookInfo, error)
	GetBookInfoByBookIds(ids []int) (*[]entities.BookInfo, error)
	DeleteBook(id int) error
}

type BookService struct {
	bookRepository    repositories.IBookRepository
	authorRepository  repositories.IAuthorRepository
	subjectRepository repositories.ISubjectRepository
	ndlApiRepository  repositories.INdlApiRepository
}

func NewBookService(bookRepository repositories.IBookRepository, authorRepository repositories.IAuthorRepository, subjectRepository repositories.ISubjectRepository, ndlApiRepository repositories.INdlApiRepository) IBookService {
	return &BookService{
		bookRepository:    bookRepository,
		authorRepository:  authorRepository,
		subjectRepository: subjectRepository,
		ndlApiRepository:  ndlApiRepository,
	}
}

func (bs *BookService) GetBooksFromNdlApi(title string, maxNum int) (*[]entities.BookInfo, error) {
	bookInfo, err := bs.ndlApiRepository.GetBooksFromNdlApi(title, maxNum)
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

func (bs *BookService) GetBooksByFuzzyTitle(title string) (*[]entities.Book, error) {
	result, err := bs.bookRepository.GetBooksByFuzzyTitle(title)
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

func (bs *BookService) GetBookInfoByISBNs(isbnSlices []int) (*[]entities.BookInfo, error) {
	result, err := bs.bookRepository.GetBookInfoByISBNs(isbnSlices)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bs *BookService) GetBookInfoByBookIds(ids []int) (*[]entities.BookInfo, error) {
	result, err := bs.bookRepository.GetBookInfoByBookIds(ids)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bs *BookService) DeleteBook(id int) error {
	err := bs.bookRepository.DeleteBook(id)
	if err != nil {
		return err
	}

	return nil
}
