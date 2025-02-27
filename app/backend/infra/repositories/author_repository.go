package repositories

import (
	"errors"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/repositories"
	"github.com/ttttai/golang/infra/models"
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
	var author []models.Author

	result := ar.db.Where("name = ?", name).Find(&author)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return models.ToAuthorDomainModels(&author), nil
}

func (ar *AuthorRepository) CreateAuthor(author *entities.Author) (*entities.Author, error) {
	authorModel := models.FromAuthorDomainModel(author)
	result := ar.db.Create(authorModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToAuthorDomainModel(authorModel), nil
}

func (ar *AuthorRepository) CreateAuthors(authors *[]entities.Author) (*[]entities.Author, error) {
	authorModels := models.FromAuthorDomainModels(authors)
	result := ar.db.Create(authorModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToAuthorDomainModels(authorModels), nil
}
