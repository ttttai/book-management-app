package repositories

import (
	"errors"

	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/infra/models"
	"gorm.io/gorm"
)

type AuthenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *AuthenticationRepository {
	return &AuthenticationRepository{
		db: db,
	}
}

func (ar *AuthenticationRepository) GetByEmail(email string) (*entities.Authentication, error) {
	var authentication models.Authentication

	result := ar.db.Where("email = ?", email).First(&authentication)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return models.ToAuthenticationDomainModel(&authentication), nil
}

func (ar *AuthenticationRepository) Create(authentication *entities.Authentication) (*entities.Authentication, error) {
	authenticationModel := models.FromAuthenticationDomainModel(authentication)

	result := ar.db.Create(authenticationModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return models.ToAuthenticationDomainModel(authenticationModel), nil
}
