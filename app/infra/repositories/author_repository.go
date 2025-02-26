package repositories

import (
	"errors"

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

func (ar *AuthorRepository) GetAuthorsByName(name string) (*[]entities.Author, error) {
	var author []entities.Author

	result := ar.db.Where("name = ?", name).Find(&author)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &author, nil
}

func (ar *AuthorRepository) CreateAuthor(author *entities.Author) (*entities.Author, error) {
	result := ar.db.Create(author)
	if result.Error != nil {
		return nil, result.Error
	}

	return author, nil
}

func (ar *AuthorRepository) CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error) {
	result := ar.db.Create(authors)
	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}
