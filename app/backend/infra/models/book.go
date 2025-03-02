package models

import (
	"time"

	"github.com/ttttai/golang/domain/entities"
)

type Book struct {
	ID                int `gorm:"primaryKey"`
	ISBN              int `gorm:"unique;not null"`
	TitleName         string
	TitleNameKana     string
	PublisherName     string
	PublisherNameKana string
	PublishDate       *string
	Price             int
	Status            int
	ReadingStartDate  *string
	ReadingEndDate    *string
	CreatedAt         time.Time `gorm:"<-:create"`
	UpdatedAt         time.Time

	Authors  []Author  `gorm:"many2many:book_authors;"`
	Subjects []Subject `gorm:"many2many:book_subjects;"`
}

func ToBookDomainModel(book *Book) *entities.Book {
	return &entities.Book{
		ID:                book.ID,
		ISBN:              book.ISBN,
		TitleName:         book.TitleName,
		TitleNameKana:     book.TitleNameKana,
		PublisherName:     book.PublisherName,
		PublisherNameKana: book.PublisherNameKana,
		PublishDate:       book.PublishDate,
		Price:             book.Price,
		Status:            book.Status,
		ReadingStartDate:  book.ReadingStartDate,
		ReadingEndDate:    book.ReadingEndDate,
	}
}

func ToBookDomainModels(books *[]Book) *[]entities.Book {
	bookEntities := make([]entities.Book, len(*books))
	for i, book := range *books {
		bookEntities[i] = *ToBookDomainModel(&book)
	}
	return &bookEntities
}

func FromBookDomainModel(book *entities.Book) *Book {
	return &Book{
		ID:                book.ID,
		ISBN:              book.ISBN,
		TitleName:         book.TitleName,
		TitleNameKana:     book.TitleNameKana,
		PublisherName:     book.PublisherName,
		PublisherNameKana: book.PublisherNameKana,
		PublishDate:       book.PublishDate,
		Price:             book.Price,
		Status:            book.Status,
		ReadingStartDate:  book.ReadingStartDate,
		ReadingEndDate:    book.ReadingEndDate,
	}
}

func FromBookDomainModels(books *[]entities.Book) *[]Book {
	bookModels := make([]Book, len(*books))
	for i, book := range *books {
		bookModels[i] = *FromBookDomainModel(&book)
	}
	return &bookModels
}
