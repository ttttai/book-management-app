package models

import (
	"time"

	"github.com/ttttai/golang/domain/entities"
)

type Author struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	NameKana  string
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time

	Books []Book `gorm:"many2many:book_authors;"`
}

func ToAuthorDomainModel(author *Author) *entities.Author {
	return &entities.Author{
		ID:       author.ID,
		Name:     author.Name,
		NameKana: author.NameKana,
	}
}

func ToAuthorDomainModels(authors *[]Author) *[]entities.Author {
	authorEntities := make([]entities.Author, len(*authors))
	for i, author := range *authors {
		authorEntities[i] = *ToAuthorDomainModel(&author)
	}
	return &authorEntities
}

func FromAuthorDomainModel(author *entities.Author) *Author {
	return &Author{
		ID:       author.ID,
		Name:     author.Name,
		NameKana: author.NameKana,
	}
}

func FromAuthorDomainModels(authors *[]entities.Author) *[]Author {
	authorModels := make([]Author, len(*authors))
	for i, author := range *authors {
		authorModels[i] = *FromAuthorDomainModel(&author)
	}
	return &authorModels
}
