package repositories

import (
	"errors"
	"fmt"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"github.com/ttttai/golang/infra/models"
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
	bookModel := models.FromBookDomainModel(book)
	result := br.db.Create(bookModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToBookDomainModel(bookModel), nil
}

func (br *BookRepository) CreateBooks(books *[]entities.Book) (*[]entities.Book, error) {
	bookModels := models.FromBookDomainModels(books)
	result := br.db.Create(bookModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToBookDomainModels(bookModels), nil
}

func (br *BookRepository) CreateBookAuthors(bookAuthors *[]entities.BookAuthor) (*[]entities.BookAuthor, error) {
	bookAuthorModels := models.FromBookAuthorDomainModels(bookAuthors)
	result := br.db.Create(bookAuthorModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToBookAuthorDomainModels(bookAuthorModels), nil
}

func (br *BookRepository) CreateBookSubjects(bookSubjects *[]entities.BookSubject) (*[]entities.BookSubject, error) {
	bookSubjectModels := models.FromBookSubjectDomainModels(bookSubjects)
	result := br.db.Create(bookSubjectModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToBookSubjectDomainModels(bookSubjectModels), nil
}

func (br *BookRepository) GetBooksByTitle(title string) (*[]entities.Book, error) {
	var book []models.Book

	result := br.db.Where("title_name = ?", title).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToBookDomainModels(&book), nil
}

func (br *BookRepository) GetBookByISBN(isbn int) (*entities.Book, error) {
	var book models.Book

	result := br.db.Where("isbn = ?", isbn).First(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return models.ToBookDomainModel(&book), nil
}

func (br *BookRepository) GetBookInfoByISBN(isbnSlices []int) (*[]entities.BookInfo, error) {
	var book []models.Book

	result := br.db.Preload("Authors").Preload("Subjects").Where("isbn IN ?", isbnSlices).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(book)

	return nil, nil
}
