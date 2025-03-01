package services

import (
	"slices"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"gorm.io/gorm"
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
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(id int) error
	UpdateBookStatus(id int, bookStatus int) (*entities.Book, error)
	GetBookInfo(title string, status []int) (*[]entities.BookInfo, error)
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

const BOOK_STATUS_NOT_PURCHASED = 0
const BOOK_STATUS_PURCHASED = 1
const BOOK_STATUS_READING = 2
const BOOK_STATUS_READ_COMPLETED = 3

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

func (bs *BookService) UpdateBook(book *entities.Book) (*entities.Book, error) {
	result, err := bs.bookRepository.UpdateBook(book)
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

func (bs *BookService) UpdateBookStatus(id int, bookStatus int) (*entities.Book, error) {
	book, err := bs.bookRepository.GetBookById(id)
	if err != nil {
		return nil, err
	}

	bookStatuses := []int{
		BOOK_STATUS_NOT_PURCHASED,
		BOOK_STATUS_PURCHASED,
		BOOK_STATUS_READING,
		BOOK_STATUS_READ_COMPLETED,
	}

	if !slices.Contains(bookStatuses, bookStatus) {
		return nil, gorm.ErrInvalidData
	}

	book.Status = bookStatus
	result, err := bs.bookRepository.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bs *BookService) GetBookInfo(title string, status []int) (*[]entities.BookInfo, error) {
	result, err := bs.bookRepository.GetBookInfo(title, status)
	if err != nil {
		return nil, err
	}

	return result, nil
}
