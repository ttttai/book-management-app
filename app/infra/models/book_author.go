package models

import "github.com/ttttai/golang/domain/entities"

type BookAuthor struct {
	BookID   int
	AuthorID int

	Book   Book   `gorm:"foreignKey:BookID;references:ID;constraint:OnDelete:CASCADE"`
	Author Author `gorm:"foreignKey:AuthorID;references:ID;constraint:OnDelete:CASCADE"`
}

func ToBookAuthorDomainModel(bookAuthor *BookAuthor) *entities.BookAuthor {
	return &entities.BookAuthor{
		BookID:   bookAuthor.BookID,
		AuthorID: bookAuthor.AuthorID,
	}
}

func ToBookAuthorDomainModels(bookAuthors *[]BookAuthor) *[]entities.BookAuthor {
	bookAuthorEntities := make([]entities.BookAuthor, len(*bookAuthors))
	for i, bookAuthor := range *bookAuthors {
		bookAuthorEntities[i] = *ToBookAuthorDomainModel(&bookAuthor)
	}
	return &bookAuthorEntities
}

func FromBookAuthorDomainModel(bookAuthor *entities.BookAuthor) *BookAuthor {
	return &BookAuthor{
		BookID:   bookAuthor.BookID,
		AuthorID: bookAuthor.AuthorID,
	}
}

func FromBookAuthorDomainModels(bookAuthors *[]entities.BookAuthor) *[]BookAuthor {
	bookAuthorModels := make([]BookAuthor, len(*bookAuthors))
	for i, bookAuthor := range *bookAuthors {
		bookAuthorModels[i] = *FromBookAuthorDomainModel(&bookAuthor)
	}
	return &bookAuthorModels
}
