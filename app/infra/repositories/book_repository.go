package repositories

import (
	"errors"
	"fmt"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) repositories.IBookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) CreateBook(book *entities.Book) (*entities.Book, error) {
	result := br.db.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (br *BookRepository) CreateBooks(books *[]entities.Book) (*[]entities.Book, error) {
	result := br.db.Create(books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (br *BookRepository) CreateBookAuthors(bookAuthors *[]entities.BookAuthor) (*[]entities.BookAuthor, error) {
	result := br.db.Create(bookAuthors)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookAuthors, nil
}

func (br *BookRepository) CreateBookSubjects(bookSubjects *[]entities.BookSubject) (*[]entities.BookSubject, error) {
	result := br.db.Create(bookSubjects)
	if result.Error != nil {
		return nil, result.Error
	}

	return bookSubjects, nil
}

func (br *BookRepository) GetBooksByTitle(title string) (*[]entities.Book, error) {
	var book []entities.Book

	result := br.db.Where("title_name = ?", title).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (br *BookRepository) GetBookByISBN(isbn int) (*entities.Book, error) {
	var book entities.Book

	result := br.db.Where("isbn = ?", isbn).First(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &book, nil
}

func (br *BookRepository) GetBookInfoByISBN(isbnSlices []int) (*[]entities.BookInfo, error) {
	var bookInfo []entities.BookInfo
	var book []entities.Book

	result := br.db.Preload("Authors").Preload("Subjects").Where("isbn IN ?", isbnSlices).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(book)

	return &bookInfo, nil
}
