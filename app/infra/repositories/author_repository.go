package repositories

import (
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) repositories.IAuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (ar *AuthorRepository) CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error) {
	result := ar.db.Create(authors)
	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}
